package models

import "time"

type QuestionDTO struct {
	ID        int         `json:"id,omitempty"`
	Text      string      `json:"text"`
	CreatedAt *time.Time  `json:"created_at,omitempty"`
	Answers   []AnswerDTO `json:"answers,omitempty"`
}

type AnswerDTO struct {
	ID         int        `json:"id,omitempty"`
	QuestionID int        `json:"question_id,omitempty"`
	UserID     string     `json:"user_id"`
	Text       string     `json:"text"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
}
