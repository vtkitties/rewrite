package internal

import (
	"fmt"
	"time"
)

/*
Define Poll model

Task: Create Poll struct with ID, Title, Options, Votes map[userID]choice, Closed flag, CreatedAt.
Acceptance: NewPoll initializes Votes (non-nil) and validates options != empty.
Hint: initialize maps with make; think about pointer vs value receivers for mutating methods.
*/

type Poll struct {
	ID        int
	Title     string
	Options   []string
	Votes     map[int]int
	Ended     bool
	CreatedAt time.Time
}

func NewPoll(id int, title string, options []string) (*Poll, error) {
	if len(options) == 0 {
		return nil, fmt.Errorf("options cannot be empty")
	}

	return &Poll{
		ID:        id,
		Title:     title,
		Options:   options,
		Votes:     make(map[int]int),
		Ended:     false,
		CreatedAt: time.Now(),
	}, nil
}

// Vote allows a user to cast a vote for a specific option in the poll
func (p *Poll) Vote(userID int, choice int) error {
	if p.Ended {
		return fmt.Errorf("poll has ended")
	}
	if choice < 0 || choice >= len(p.Options) {
		return fmt.Errorf("invalid choice")
	}
	p.Votes[userID] = choice
	return nil
}
