package usecase

import (
	"github.com/haikalvidya/go-simple-todo/internal/delivery/payload"
	"github.com/haikalvidya/go-simple-todo/internal/models"
	"gorm.io/gorm"
)

type IActivityUsecase interface {
	GetAllActivities() ([]*payload.ActivityInfo, error)
	SelectActivityByID(id int64) (*payload.ActivityInfo, error)
	CreateActivity(activity *payload.CreateActivityRequest) (*payload.ActivityInfo, error)
	UpdateActivity(id int64, activity *payload.UpdateActivityRequest) (*payload.ActivityInfo, error)
	DeleteActivity(id int64) error
}

type activityUsecase usecaseType

func (u *activityUsecase) GetAllActivities() ([]*payload.ActivityInfo, error) {
	activities, err := u.Repo.Activity.GetAll()
	if err != nil {
		return nil, err
	}

	activityInfos := []*payload.ActivityInfo{}
	for _, activity := range activities {
		activityInfos = append(activityInfos, activity.PublicInfo())
	}

	return activityInfos, nil
}

func (u *activityUsecase) SelectActivityByID(id int64) (*payload.ActivityInfo, error) {
	activity, err := u.Repo.Activity.SelectByID(id)
	if err != nil {
		return nil, err
	}

	return activity.PublicInfo(), nil
}

func (u *activityUsecase) CreateActivity(activity *payload.CreateActivityRequest) (*payload.ActivityInfo, error) {
	activityModel := &models.Activity{
		Title: activity.Title,
		Email: activity.Email,
	}

	var err error

	err = u.Repo.Tx.DoInTransaction(func(tx *gorm.DB) error {
		activityModel, err = u.Repo.Activity.CreateTx(tx, activityModel)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return activityModel.PublicInfo(), nil
}

func (u *activityUsecase) UpdateActivity(id int64, activity *payload.UpdateActivityRequest) (*payload.ActivityInfo, error) {
	activityModel := &models.Activity{
		ID:    id,
		Title: activity.Title,
		Email: activity.Email,
	}

	var err error

	err = u.Repo.Tx.DoInTransaction(func(tx *gorm.DB) error {
		err = u.Repo.Activity.UpdateTx(tx, activityModel)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return activityModel.PublicInfo(), nil
}

func (u *activityUsecase) DeleteActivity(id int64) error {
	activity, err := u.Repo.Activity.SelectByID(id)
	if err != nil {
		return err
	}

	err = u.Repo.Tx.DoInTransaction(func(tx *gorm.DB) error {
		err = u.Repo.Activity.DeleteTx(tx, activity)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
