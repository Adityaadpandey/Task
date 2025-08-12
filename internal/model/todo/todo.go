package todo

import (
	"time"

	"github.com/adityaadpandey/go-boilerplate/internal/model"
	"github.com/adityaadpandey/go-boilerplate/internal/model/category"
	"github.com/adityaadpandey/go-boilerplate/internal/model/comment"
	"github.com/google/uuid"
)

type Status string

const (
	StatusDraft     Status = "draft"
	StatusActive    Status = "active"
	StatusCompleted Status = "completed"
	StatusArchived  Status = "archived"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

type Todo struct {
	model.Base
	UserID       string     `json:"userId" db:"user_id"`
	Title        string     `json:"title" db:"title"`
	Description  *string    `json:"description" db:"description"`
	Status       Status     `json:"status" db:"status"`
	Priority     Priority   `json:"priority" db:"priority"`
	DueDate      *time.Time `json:"dueDate" db:"due_date"`
	CompletdeAt  *time.Time `json:"completedAt,omitempty" db:"completed_at"`
	ParentTodoID *uuid.UUID `json:"parentTodoId,omitempty" db:"parent_todo_id"`
	CategoryID   *uuid.UUID `json:"categoryId,omitempty" db:"category_id"`
	Metadata     *Metadata  `json:"metadata,omitempty" db:"metadata"`
	SordOrder    int        `json:"sordOrder" db:"sord_order"`
}

type Metadata struct {
	Tags      []string `json:"tags"`
	Reminder  *string  `json:"reminder"`
	Color     *string  `json:"color"`
	Difficlty *string  `json:"difficulty"`
}

type PopulatedTodo struct {
	Todo
	Category *category.Category `json:"category" db:"category"`
	Children []Todo             `json:"children" db:"children"`
	Comments []comment.Comment  `json:"comments" db:"comments"`
}

type TodoStats struct {
	Total     int `json:"total"`
	Draft     int `json:"draft"`
	Active    int `json:"active"`
	Completed int `json:"completed"`
	Archived  int `json:"archived"`
	Overdue   int `json:"overdue"`
}

func (t *Todo) IsOverdue() bool {
	return t.DueDate != nil && t.DueDate.Before(time.Now()) && t.Status != StatusCompleted
}

func (t *Todo) CanHaveChildren() bool {
	return t.ParentTodoID == nil
}
