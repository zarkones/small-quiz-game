package main

import (
	"bufio"
	"bytes"
	"contracts"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	USERNAME string
	API_HOST string
)

var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Quiz game of guessing the right answers",
	Run: func(cmd *cobra.Command, args []string) {
		if len(USERNAME) == 0 {
			fmt.Println("username not provided")
			os.Exit(1)
		}
		if len(API_HOST) == 0 {
			fmt.Println("api host not provided")
			os.Exit(1)
		}

		resp, err := http.DefaultClient.Get(API_HOST + "/v1/questions")
		if err != nil {
			fmt.Println("failed to get questions:", err)
			os.Exit(1)
		}

		var questions contracts.GetQuestionsRespCtx
		answers := contracts.AnswerQuestionsReqCtx{
			Username:           USERNAME,
			QuestionsAndAnswer: map[string]string{},
		}

		if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
			fmt.Println("failed to parse server's response:", err)
			os.Exit(1)
		}

		for i, q := range questions {
			fmt.Println("Question #"+strconv.Itoa(i+1)+":", q.Question)
			fmt.Println()
			fmt.Println("A:", q.Answers.A)
			fmt.Println("B:", q.Answers.B)
			fmt.Println("C:", q.Answers.C)
			fmt.Println("D:", q.Answers.D)
			fmt.Println()

			reader := bufio.NewReader(os.Stdin)

			// We could use a for statement to handle this.
			// However, for this really specific case it's
			// my opinion that a goto statement is clear
			// and simple enough, with minimal risk to create
			// unexpected behavior.
		GET_USER_ANSWER:
			fmt.Print("_> ")
			rawInput, _, err := reader.ReadLine()
			if err != nil {
				fmt.Println("error while handling user input:", err)
				os.Exit(1)
			}
			input := strings.ToUpper(string(rawInput))
			if !slices.Contains([]string{"A", "B", "C", "D"}, input) {
				fmt.Println("invalid answer:", input)
				goto GET_USER_ANSWER
			}
			answers.QuestionsAndAnswer[q.Question] = input
		}

		serializedAnswers, err := json.Marshal(&answers)
		if err != nil {
			fmt.Println("failed to serialize user's answers:", err)
			os.Exit(1)
		}

		req, err := http.NewRequest(http.MethodPut, API_HOST+"/v1/questions", bytes.NewReader(serializedAnswers))
		if err != nil {
			fmt.Println("failed to send user's answers:", err)
			os.Exit(1)
		}

		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("failed to read server's response to user's answers:", err)
			os.Exit(1)
		}

		var scoreCtx contracts.AnswerQuestionsRespCtx
		if err := json.NewDecoder(resp.Body).Decode(&scoreCtx); err != nil {
			fmt.Println("failed to deserialize server's response:", err)
			os.Exit(1)
		}

		fmt.Println("Score:", scoreCtx.Score)
		fmt.Println("You're better than", scoreCtx.BetterThanPercent, "percentage of players")
	},
}

func main() {
	rootCmd.PersistentFlags().StringVar(&USERNAME, "name", "", "Player's username")
	rootCmd.PersistentFlags().StringVar(&API_HOST, "host", "", "API host")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
