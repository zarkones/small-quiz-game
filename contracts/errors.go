package contracts

import "errors"

var (
	ErrInvalidUsername    = errors.New("invalid username")
	ErrInvalidQnaLen      = errors.New("invalid questions and answers map length")
	ErrInvalidQuestionLen = errors.New("invalid question length")
	ErrInvalidAnswerLen   = errors.New("invalid answer length")
)
