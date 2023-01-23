package repository

import (
	"github.com/haikalvidya/go-simple-todo/internal/models"
	"gorm.io/gorm"
)

type IActivityRepository interface {
	GetAll() ([]*models.Activity, error)
	SelectByID(id int64) (*models.Activity, error)
	CreateTx(tx *gorm.DB, activity *models.Activity) (*models.Activity, error)
	DeleteTx(tx *gorm.DB, activity *models.Activity) error
	UpdateTx(tx *gorm.DB, activity *models.Activity) error
}

type activityRepository repositoryType

func (r *activityRepository) GetAll() ([]*models.Activity, error) {
	activities := []*models.Activity{}
	err := r.DB.Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *activityRepository) SelectByID(id int64) (*models.Activity, error) {
	activity := &models.Activity{}
	err := r.DB.Where("id = ?", id).First(activity).Error
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (r *activityRepository) CreateTx(tx *gorm.DB, activity *models.Activity) (*models.Activity, error) {
	err := tx.Create(activity).Error
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (r *activityRepository) DeleteTx(tx *gorm.DB, activity *models.Activity) error {
	err := tx.Delete(activity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *activityRepository) UpdateTx(tx *gorm.DB, activity *models.Activity) error {
	err := tx.Save(activity).Error
	if err != nil {
		return err
	}
	return nil
}
