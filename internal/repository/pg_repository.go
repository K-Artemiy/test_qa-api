package repository

import (
	"errors"

	"gorm.io/gorm"

	"test_qa-api/internal/models"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateQuestion(q *models.Question) error {
	return r.db.Raw(
		`INSERT INTO questions(text) 
		VALUES (?) 
		RETURNING id, text, created_at`,
		q.Text,
	).Scan(q).Error
}

func (r *Repo) ListQuestions() ([]models.Question, error) {
	var qs []models.Question
	err := r.db.Raw(
		`SELECT id, text, created_at 
		FROM questions 
		ORDER BY id`,
	).Scan(&qs).Error
	return qs, err
}

func (r *Repo) GetQuestion(id int) (*models.Question, []models.Answer, error) {
	var q models.Question
	if err := r.db.Raw(
		`SELECT id, text, created_at 
		FROM questions 
		WHERE id = ?`,
		id,
	).Scan(&q).Error; err != nil {
		return nil, nil, err
	}
	if q.ID == 0 {
		return nil, nil, gorm.ErrRecordNotFound
	}
	var ans []models.Answer
	if err := r.db.Raw(
		`SELECT id, question_id, user_id, text, created_at 
		FROM answers 
		WHERE question_id = ? 
		ORDER BY id`,
		id,
	).Scan(&ans).Error; err != nil {
		return nil, nil, err
	}
	return &q, ans, nil
}

func (r *Repo) DeleteQuestion(id int) error {
	res := r.db.Exec(
		`DELETE FROM questions 
		WHERE id = ?`,
		id,
	)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *Repo) CreateAnswer(a *models.Answer) error {
	var exists bool
	if err := r.db.Raw(
		`SELECT true 
		FROM questions 
		WHERE id = ?`,
		a.QuestionID,
	).Scan(&exists).Error; err != nil {
		return err
	}
	if !exists {
		return errors.New("question not found")
	}
	return r.db.Raw(
		`INSERT INTO answers(question_id, user_id, text) 
		VALUES (?, ?, ?) 
		RETURNING id, question_id, user_id, text, created_at`,
		a.QuestionID,
		a.UserID,
		a.Text,
	).Scan(a).Error
}

func (r *Repo) GetAnswer(id int) (*models.Answer, error) {
	var a models.Answer
	if err := r.db.Raw(
		`SELECT id, question_id, user_id, text, created_at 
		FROM answers 
		WHERE id = ?`,
		id,
	).Scan(&a).Error; err != nil {
		return nil, err
	}
	if a.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &a, nil
}

func (r *Repo) DeleteAnswer(id int) error {
	res := r.db.Exec(
		`DELETE FROM answers 
		WHERE id = ?`,
		id,
	)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
