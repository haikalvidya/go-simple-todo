package models

import (
	"time"

	"github.com/haikalvidya/go-simple-todo/internal/delivery/payload"
	"gorm.io/gorm"
)

type Activity struct {
	ID        int64          `db:"id"`
	Title     string         `db:"title"`
	Email     string         `db:"email"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at"`
}

func (Activity) TableName() string {
	return "activities"
}

func (a *Activity) BeforeCreate(tx *gorm.DB) (err error) {
	a.CreatedAt = time.Now()
	return
}

func (a *Activity) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	a.UpdatedAt = now
	return
}

func (a *Activity) PublicInfo() *payload.ActivityInfo {
	res := &payload.ActivityInfo{
		ID:        a.ID,
		Title:     a.Title,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}

	if a.Email != "" {
		res.Email = &a.Email
	} else {
		res.Email = nil
	}

	if a.DeletedAt.Valid {
		res.DeletedAt = &a.DeletedAt.Time
	} else {
		res.DeletedAt = nil
	}

	return res
}
