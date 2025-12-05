package repository

import "test_qa-api/internal/models"

type Repo interface {
	CreateQuestion(q *models.Question) error
	ListQuestions() ([]models.Question, error)
	GetQuestion(id int) (*models.Question, []models.Answer, error)
	DeleteQuestion(id int) error
	CreateAnswer(a *models.Answer) error
	GetAnswer(id int) (*models.Answer, error)
	DeleteAnswer(id int) error
}
