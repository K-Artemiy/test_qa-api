package services

import (
	"test_qa-api/internal/models"
)

func convQuestionToDTO(q *models.Question, answers []models.Answer) *models.QuestionDTO {
	dto := &models.QuestionDTO{
		ID:        q.ID,
		Text:      q.Text,
		CreatedAt: &q.CreatedAt,
	}

	if len(answers) > 0 {
		dto.Answers = make([]models.AnswerDTO, 0, len(answers))
		for _, a := range answers {
			dto.Answers = append(dto.Answers, *convAnswerToDTO(&a))
		}
	}
	return dto
}

func convAnswerToDTO(a *models.Answer) *models.AnswerDTO {
	dto := &models.AnswerDTO{
		ID:         a.ID,
		QuestionID: a.QuestionID,
		UserID:     a.UserID,
		Text:       a.Text,
		CreatedAt:  &a.CreatedAt,
	}
	return dto
}

func convDTOToQuestion(q *models.QuestionDTO) *models.Question {
	return &models.Question{
		Text: q.Text,
	}
}

func convDTOToAnswer(a *models.AnswerDTO, questionID int) *models.Answer {
	return &models.Answer{
		QuestionID: questionID,
		UserID:     a.UserID,
		Text:       a.Text,
	}
}
