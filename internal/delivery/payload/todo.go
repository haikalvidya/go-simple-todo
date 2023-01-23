package payload

import "time"

type TodoInfo struct {
	ID              int64     `json:"id"`
	Title           string    `json:"title"`
	ActivityGroupId int64     `json:"activity_group_id"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	// can be null
	DeletedAt *time.Time `json:"deletedAt"`
}

type TodoListResponseWithoutDeletedAt struct {
	ID              int64     `json:"id"`
	Title           string    `json:"title"`
	ActivityGroupId int64     `json:"activity_group_id"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type CreateTodoRequest struct {
	Title           string `json:"title" validate:"required"`
	ActivityGroupId int64  `json:"activity_group_id" validate:"required"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}

type UpdateTodoRequest struct {
	Title    string `json:"title"`
	IsActive bool   `json:"is_active"`
	Priority string `json:"priority"`
}

const (
	TODO_ID_NOT_FOUND = "Todo with ID %v Not Found"
)
