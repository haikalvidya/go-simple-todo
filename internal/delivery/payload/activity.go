package payload

import "time"

type ActivityInfo struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Email     *string   `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// can be null
	DeletedAt *time.Time `json:"deletedAt"`
}

type ActivityListResponseWithoutDeletedAt struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Email     *string   `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateActivityRequest struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email"`
}

type UpdateActivityRequest struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email"`
}

const (
	ERROR_ACTIVITY_NOT_FOUND = "Activity with ID %d Not Found"
)
