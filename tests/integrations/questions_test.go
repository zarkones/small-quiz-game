package ctrl

import (
	"api/core"
	"bytes"
	"contracts"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestQuestionsAndAnswers(t *testing.T) {
	is := is.New(t)

	mux := http.NewServeMux()

	core.InitRoutes(mux)

	srv := &http.Server{
		Handler: mux,
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/questions", nil)
	w := httptest.NewRecorder()

	srv.Handler.ServeHTTP(w, req)

	res := w.Result()

	is.Equal(res.StatusCode, http.StatusOK)

	var questions contracts.GetQuestionsRespCtx
	is.NoErr(json.NewDecoder(res.Body).Decode(&questions))

	if len(questions) == 0 {
		t.Log("questions have length of zero")
		t.FailNow()
	}

	username := "user123"

	answers := contracts.AnswerQuestionsReqCtx{
		Username:           username,
		QuestionsAndAnswer: map[string]string{},
	}

	for _, q := range questions {
		if len(q.Question) == 0 {
			t.Log("invalid question length")
			t.FailNow()
		}
		if len(q.Answers.A) == 0 {
			t.Log("invalid potential answer A length")
			t.FailNow()
		}
		if len(q.Answers.B) == 0 {
			t.Log("invalid potential answer B length")
			t.FailNow()
		}
		if len(q.Answers.C) == 0 {
			t.Log("invalid potential answer C length")
			t.FailNow()
		}
		if len(q.Answers.D) == 0 {
			t.Log("invalid potential answer D length")
			t.FailNow()
		}

		answers.QuestionsAndAnswer[q.Question] = "A"
	}

	answersJson, err := json.Marshal(&answers)
	is.NoErr(err)

	req = httptest.NewRequest(http.MethodPut, "/v1/questions", bytes.NewReader(answersJson))
	w = httptest.NewRecorder()

	srv.Handler.ServeHTTP(w, req)

	res = w.Result()

	is.Equal(res.StatusCode, http.StatusOK)

	var respCtx contracts.AnswerQuestionsRespCtx

	is.NoErr(json.NewDecoder(res.Body).Decode(&respCtx))

	if respCtx.BetterThanPercent != 100 {
		t.Log("Unexpected Relative Score:", respCtx.BetterThanPercent)
		t.FailNow()
	}

	t.Log("Score:", respCtx.Score)
	t.Log("Relative Score:", respCtx.BetterThanPercent)
}
