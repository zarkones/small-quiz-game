package data

import (
	"testing"

	"github.com/matryer/is"
)

func TestScore(t *testing.T) {
	is := is.New(t)

	user := "user123"
	score := 4

	InsertScore(user, score)

	storedScore := scores[user]
	if storedScore != score {
		t.Log("stored score isn't the user score")
		t.FailNow()
	}

	score = 5
	InsertScore(user, score)

	storedScore = scores[user]
	if storedScore != score {
		t.Log("stored score isn't the updated user score")
		t.FailNow()
	}

	InsertScore("user1", 1)
	InsertScore("user2", 1)
	InsertScore("user3", 1)
	InsertScore("user4", 1)
	InsertScore("user5", 1)
	InsertScore("user6", 10)
	InsertScore("user7", 10)
	InsertScore("user8", 10)
	InsertScore("user9", 10)
	InsertScore("user10", 10)

	perc, err := CalcRelativeScore(user)
	is.NoErr(err)

	if perc != 50 {
		t.Log("'better than' percentage calculated wrong:", perc)
		t.Log("User score:", score)
		t.Log("Amount of players:", len(scores))
		t.FailNow()
	}

	t.Log("Better than:", perc, "%")

	InsertScore("user11", 10)
	InsertScore("user12", 10)
	InsertScore("user13", 10)
	InsertScore("user14", 10)
	InsertScore("user15", 10)
	InsertScore("user16", 10)
	InsertScore("user17", 10)
	InsertScore("user18", 10)
	InsertScore("user19", 10)
	InsertScore("user20", 10)

	perc, err = CalcRelativeScore(user)
	is.NoErr(err)

	if perc != 25 {
		t.Log("'better than' percentage calculated wrong:", perc)
		t.Log("User score:", score)
		t.Log("Amount of players:", len(scores))
		t.FailNow()
	}

	t.Log("Better than:", perc, "%")
}
