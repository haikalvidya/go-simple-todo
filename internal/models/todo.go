package models

import (
	"time"

	"github.com/haikalvidya/go-simple-todo/internal/delivery/payload"
	"gorm.io/gorm"
)

type Todo struct {
	ID              int64          `db:"id"`
	Title           string         `db:"title"`
	ActivityGroupId int64          `db:"activity_group_id"`
	IsActive        bool           `db:"is_active"`
	Priority        string         `db:"priority"`
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
	DeletedAt       gorm.DeletedAt `db:"deleted_at"`
}

func (Todo) TableName() string {
	return "todos"
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now()
	if t.Priority == "" {
		t.Priority = "very-high"
	}
	return
}

func (t *Todo) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	t.UpdatedAt = now
	return
}

func (t *Todo) PublicInfo() *payload.TodoInfo {
	res := &payload.TodoInfo{
		ID:              t.ID,
		Title:           t.Title,
		ActivityGroupId: t.ActivityGroupId,
		IsActive:        t.IsActive,
		Priority:        t.Priority,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
	}

	if t.DeletedAt.Valid {
		res.DeletedAt = &t.DeletedAt.Time
	} else {
		res.DeletedAt = nil
	}

	return res
}
