package contracts

type PossibleAnswers struct {
	A string
	B string
	C string
	D string
}

type Question struct {
	Question string
	Answers  PossibleAnswers
}

type AnsweredQuestion struct {
	Question
	CorrectAnswer string
}

type GetQuestionsRespCtx []Question

type AnswerQuestionsReqCtx struct {
	Username string
	// QuestionsAndAnswer is a map[question]a|b|c|d
	QuestionsAndAnswer map[string]string
}

func (ctx *AnswerQuestionsReqCtx) Validate() (err error) {
	if len(ctx.Username) == 0 {
		return ErrInvalidUsername
	}
	if len(ctx.QuestionsAndAnswer) == 0 {
		return ErrInvalidQnaLen
	}
	for q, a := range ctx.QuestionsAndAnswer {
		if len(q) == 0 {
			return ErrInvalidQuestionLen
		}
		switch a {
		case "A", "B", "C", "D":
			continue
		default:
			return ErrInvalidAnswerLen
		}
	}

	return nil
}

type AnswerQuestionsRespCtx struct {
	Score             int
	BetterThanPercent float32
}
