package domain

import "time"

// Vote value data.
type Vote struct {
	Option    VoteOption
	Username  string
	Email     string
	CreatedAt time.Time
}

// Vote Option data.
type VoteOption struct {
	ID          int
	Description string
}
