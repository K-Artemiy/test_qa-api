package models

import "time"

type Question struct {
	ID        int
	Text      string
	CreatedAt time.Time
}

type Answer struct {
	ID         int
	QuestionID int
	UserID     string
	Text       string
	CreatedAt  time.Time
}
