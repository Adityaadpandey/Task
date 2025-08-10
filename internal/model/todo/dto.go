package todo

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// -----------------------------------------------------------------

type CreateTodoPayload struct {
	Title        string     `json:"title" validate:"required,min=1,max=255"`
	Description  *string    `json:"description,omitempty" validate:"omitempty,max=1000"`
	Priority     *Priority  `json:"priority" validate:"required,oneof=low medium high"`
	DueDate      *time.Time `json:"dueDate"`
	ParentTodoID *uuid.UUID `json:"parentTodoId" validate:"omitempty,uuid"`
	CategoryID   *uuid.UUID `json:"categoryId" validate:"omitempty,uuid"`
	Metadata     *Metadata  `json:"metadta"`
}

func (p *CreateTodoPayload) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

// -----------------------------------------------------------------

type UpdateTodoPayload struct {
	ID           uuid.UUID  `param:"id" validate:"required,uuid"`
	Title        *string    `json:"title" validate:"required,min=1,max=255"`
	Description  *string    `json:"description,omitempty" validate:"omitempty,max=1000"`
	Status       *Status    `json:"status" validate:"required,oneof=draft active completed archived"`
	Priority     *Priority  `json:"priority" validate:"required,oneof=low medium high"`
	DueDate      *time.Time `json:"dueDate"`
	ParentTodoID *uuid.UUID `json:"parentTodoId" validate:"omitempty,uuid"`
	CategoryID   *uuid.UUID `json:"categoryId" validate:"omitempty,uuid"`
	Metadata     *Metadata  `json:"metadta"`
}

func (p *UpdateTodoPayload) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

// -----------------------------------------------------------------

type GetTodosQuery struct {
	Page         *int       `query:"page" validate:"omitempty,min=1"`
	Limit        *int       `query:"limit" validate:"omitempty,min=1,max=100"`
	Sort         *string    `query:"sort" validate:"omitempty,oneof=created_at updated_at title priority due_date status"`
	Order        *string    `query:"order" validate:"omitempty,oneof=asc desc"`
	Search       *string    `query:"search" validate:"omitempty,min=1"`
	Status       *Status    `query:"status" validate:"omitempty,oneof=draft active completed archived"`
	Priority     *Priority  `query:"priority" validate:"omitempty,oneof=low medium high"`
	CategoryID   *uuid.UUID `query:"categoryId" validate:"omitempty,uuid"`
	ParentTodoID *uuid.UUID `query:"parentTodoId" validate:"omitempty,uuid"`
	DueFrom      *time.Time `query:"dueFrom"`
	DueTo        *time.Time `query:"dueTo"`
	Overdue      *bool      `query:"overdue"`
	Completed    *bool      `query:"completed"`
}

func (q *GetTodosQuery) Validate() error {
	validate := validator.New()

	if err := validate.Struct(q); err != nil {
		return err
	}

	if q.Page == nil {
		defaltPage := 1
		q.Page = &defaltPage
	}
	if q.Limit == nil {
		defaultLimit := 20
		q.Limit = &defaultLimit
	}
	if q.Sort == nil {
		defaultSort := "created_at"
		q.Sort = &defaultSort
	}
	if q.Order == nil {
		defaultOrder := "desc"
		q.Order = &defaultOrder
	}

	return nil
}

// -----------------------------------------------------------------

type GetTodoByIDPayload struct {
	ID uuid.UUID `param:"id" validate:"required,uuid"`
}

func (p *GetTodoByIDPayload) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

// -----------------------------------------------------------------

type DeleteTodoPayload struct {
	ID uuid.UUID `param:"id" validate:"required,uuid"`
}

func (p *DeleteTodoPayload) Validate() error {
	validator := validator.New()
	return validator.Struct(p)
}

// -----------------------------------------------------------------

type GetTodoStatsPayload struct{}

func (p *GetTodoStatsPayload) Validate() error {
	validator := validator.New()
	return validator.Struct(p)
}
