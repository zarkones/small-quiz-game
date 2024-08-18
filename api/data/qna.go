package data

import "contracts"

var QnA = []contracts.AnsweredQuestion{
	{
		Question: contracts.Question{
			Question: "Which is the best programming language?",
			Answers: contracts.PossibleAnswers{
				A: "Java",
				B: "JavaScript",
				C: "Go",
				D: "Haxe",
			},
		},
		CorrectAnswer: "C",
	},
	{
		Question: contracts.Question{
			Question: "What's the meaning of life?",
			Answers: contracts.PossibleAnswers{
				A: "Being a good person",
				B: "Fast cars",
				C: "Adventure",
				D: "42",
			},
		},
		CorrectAnswer: "D",
	},
	{
		Question: contracts.Question{
			Question: "Who is the true hero?",
			Answers: contracts.PossibleAnswers{
				A: "Superman",
				B: "Samwise the Brave",
				C: "Batman",
				D: "Frodo",
			},
		},
		CorrectAnswer: "B",
	},
	{
		Question: contracts.Question{
			Question: "Which is the most used sentence of this decade?",
			Answers: contracts.PossibleAnswers{
				A: "It worked for me locally",
				B: "We're doing deployment on Friday",
				C: "I as an AI language model",
				D: "Like, share, and subscribe",
			},
		},
		CorrectAnswer: "C",
	},
	{
		Question: contracts.Question{
			Question: "Which will be the most used sentence of the year 2050?",
			Answers: contracts.PossibleAnswers{
				A: "Do you need the real me for this conversation",
				B: "ClosedAI just announced ChatGPT-7843, I'm sure that ChatGPT-7844 will be an AGI",
				C: "Sorry I'm late, boss. My car had an automatic software update",
				D: "It worked for me locally",
			},
		},
		CorrectAnswer: "A",
	},
}
