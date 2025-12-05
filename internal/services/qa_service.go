package services

import (
	"errors"

	"test_qa-api/internal/models"
	"test_qa-api/internal/repository"
)

type QAService struct {
	repo repository.Repo
}

func NewQAService(r repository.Repo) *QAService {
	return &QAService{repo: r}
}

func (s *QAService) CreateQuestion(dto *models.QuestionDTO) (*models.QuestionDTO, error) {
	if dto.Text == "" {
		return nil, errors.New("text is required")
	}
	q := convDTOToQuestion(dto)
	if err := s.repo.CreateQuestion(q); err != nil {
		return nil, err
	}
	return convQuestionToDTO(q, nil), nil
}

func (s *QAService) ListQuestions() ([]models.QuestionDTO, error) {
	qs, err := s.repo.ListQuestions()
	if err != nil {
		return nil, err
	}
	out := make([]models.QuestionDTO, 0, len(qs))
	for _, q := range qs {
		dto := convQuestionToDTO(&q, nil)
		out = append(out, *dto)
	}
	return out, nil
}

func (s *QAService) GetQuestion(id int) (*models.QuestionDTO, error) {
	q, answers, err := s.repo.GetQuestion(id)
	if err != nil {
		return nil, err
	}
	return convQuestionToDTO(q, answers), nil
}

func (s *QAService) DeleteQuestion(id int) error {
	return s.repo.DeleteQuestion(id)
}

func (s *QAService) CreateAnswer(questionID int, dto *models.AnswerDTO) (*models.AnswerDTO, error) {
	if dto.UserID == "" || dto.Text == "" {
		return nil, errors.New("user_id and text are required")
	}
	a := convDTOToAnswer(dto, questionID)
	if err := s.repo.CreateAnswer(a); err != nil {
		return nil, err
	}
	return convAnswerToDTO(a), nil
}

func (s *QAService) GetAnswer(id int) (*models.AnswerDTO, error) {
	a, err := s.repo.GetAnswer(id)
	if err != nil {
		return nil, err
	}
	return convAnswerToDTO(a), nil
}

func (s *QAService) DeleteAnswer(id int) error {
	return s.repo.DeleteAnswer(id)
}
