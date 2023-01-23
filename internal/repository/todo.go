package repository

import (
	"github.com/haikalvidya/go-simple-todo/internal/models"
	"gorm.io/gorm"
)

type ITodoRepository interface {
	GetAll() ([]*models.Todo, error)
	SelectByID(id int64) (*models.Todo, error)
	CreateTx(tx *gorm.DB, todo *models.Todo) (*models.Todo, error)
	DeleteTx(tx *gorm.DB, todo *models.Todo) error
	UpdateTx(tx *gorm.DB, todo *models.Todo) error
	SelectByQueryActivityGroupID(activityGroupID int64) ([]*models.Todo, error)
}

type todoRepository repositoryType

func (r *todoRepository) GetAll() ([]*models.Todo, error) {
	todos := []*models.Todo{}
	err := r.DB.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) SelectByID(id int64) (*models.Todo, error) {
	todo := &models.Todo{}
	err := r.DB.Where("id = ?", id).First(todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *todoRepository) CreateTx(tx *gorm.DB, todo *models.Todo) (*models.Todo, error) {
	err := tx.Create(todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *todoRepository) DeleteTx(tx *gorm.DB, todo *models.Todo) error {
	err := tx.Delete(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *todoRepository) UpdateTx(tx *gorm.DB, todo *models.Todo) error {
	err := tx.Save(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *todoRepository) SelectByQueryActivityGroupID(activityGroupID int64) ([]*models.Todo, error) {
	todos := []*models.Todo{}
	err := r.DB.Where("activity_group_id = ?", activityGroupID).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}
