package entity

import "time"

// Vote value data.
type Vote struct {
	PollID        int       `json:"pollId"`
	AlternativeID int       `json:"alternativeId"`
	Username      string    `json:"username"`
	CreatedAt     time.Time `json:"createdAt"`
}
