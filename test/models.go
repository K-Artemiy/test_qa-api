package test

type QuestionResponse struct {
	ID        int              `json:"id"`
	Text      string           `json:"text"`
	CreatedAt string           `json:"created_at"`
	Answers   []AnswerResponse `json:"answers,omitempty"`
}

type CreateQuestionRequest struct {
	Text string `json:"text"`
}

type AnswerResponse struct {
	ID         int    `json:"id"`
	QuestionID int    `json:"question_id"`
	UserID     string `json:"user_id"`
	Text       string `json:"text"`
	CreatedAt  string `json:"created_at"`
}

type CreateAnswerRequest struct {
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}
