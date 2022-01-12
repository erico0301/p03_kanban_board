package modeltask

import (
	"p03_kanban_board/model/modeluser"
	"time"
)

type Request struct {
	ID          uint64 `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status,omitempty"`
	CategoryID  uint   `json:"category_id"`
	UserID      uint   `json:"user_id,omitempty"`
}

type RequestUpdate struct {
	ID          uint64 `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id,omitempty"`
}

type RequestUpdateStatus struct {
	ID     uint64 `json:"id,omitempty"`
	Status *bool  `json:"status"`
	UserID uint   `json:"user_id,omitempty"`
}

type RequestUpdateCategory struct {
	ID         uint64 `json:"id,omitempty"`
	CategoryID uint   `json:"category_id"`
	UserID     uint   `json:"user_id,omitempty"`
}

type ResponseStore struct {
	ID          uint64     `json:"id"`
	Title       string     `json:"title"`
	Status      bool       `json:"status"`
	Description string     `json:"description"`
	UserID      uint       `json:"user_id"`
	CategoryID  uint       `json:"category_id"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type ResponseGet struct {
	ResponseStore
	User modeluser.Response `json:"user"`
}

type ExampleResponseDelete struct {
	Message string `json:"message" example:"Task has been deleted"`
}
