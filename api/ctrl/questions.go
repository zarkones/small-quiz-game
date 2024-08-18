package ctrl

import (
	"api/data"
	"contracts"
	"encoding/json"
	"net/http"
)

// GetQuestions returns a questions without provided answers.
func GetQuestions() http.HandlerFunc {
	var questions = make(contracts.GetQuestionsRespCtx, len(data.QnA))

	// Pre-load questions in a format suitable for the consuming clients.
	for i, answeredQuestion := range data.QnA {
		questions[i] = contracts.Question{
			Question: answeredQuestion.Question.Question,
			Answers:  answeredQuestion.Answers,
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(&questions); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Answer takes responses to questions from a user.
func Answer(w http.ResponseWriter, r *http.Request) {
	var ctx contracts.AnswerQuestionsReqCtx

	if err := json.NewDecoder(r.Body).Decode(&ctx); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := ctx.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	score := 0

	for _, qna := range data.QnA {
		answer, ok := ctx.QuestionsAndAnswer[qna.Question.Question]
		// Abort the request if a user hasn't provided an answer any of the existing questions.
		if !ok {
			http.Error(w, "no answer provided to question: "+qna.Question.Question, http.StatusUnprocessableEntity)
			return
		}

		if answer != qna.CorrectAnswer {
			continue
		}

		score += 1
	}

	data.InsertScore(ctx.Username, score)

	betterThanPerc, err := data.CalcRelativeScore(ctx.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respCtx := contracts.AnswerQuestionsRespCtx{
		Score:             score,
		BetterThanPercent: betterThanPerc,
	}

	if err := json.NewEncoder(w).Encode(&respCtx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
