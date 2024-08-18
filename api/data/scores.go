package data

import (
	"errors"
	"sync"
)

// scores is a map of usernames and their scores.
var scores = map[string]int{}

var scoresMux = &sync.Mutex{}

func InsertScore(username string, score int) {
	defer scoresMux.Unlock()
	scoresMux.Lock()

	scores[username] = score
}

func CalcRelativeScore(username string) (percentage float32, err error) {
	defer scoresMux.Unlock()
	scoresMux.Lock()

	playerScore, ok := scores[username]
	if !ok {
		return 0, ErrUserNotFound
	}

	if len(scores) == 1 {
		// The user is the only player.
		return 100, nil
	}

	betterPerforming := float32(0)
	worsePerforming := float32(0)

	for currUsername, score := range scores {
		if currUsername == username {
			continue
		}

		if score > playerScore {
			betterPerforming += 1
		}
		if score < playerScore {
			worsePerforming += 1
		}
	}

	decimalPerc := worsePerforming / float32(len(scores)-1)

	return decimalPerc * 100, nil
}

var (
	ErrUserNotFound = errors.New("user of that name not found")
)
