package handler

import (
	"github.com/adityaadpandey/go-boilerplate/internal/server"
	"github.com/adityaadpandey/go-boilerplate/internal/service"
)

type Handlers struct {
	Health          *HealthHandler
	OpenAPI         *OpenAPIHandler
	TodoHandler     *TodoHandler
	CommentHandler  *CommentHandler
	CategoryHandler *CategoryHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health:          NewHealthHandler(s),
		OpenAPI:         NewOpenAPIHandler(s),
		TodoHandler:     NewTodoHandler(s, services.Todo),
		CommentHandler:  NewCommentHandler(s, services.Comment),
		CategoryHandler: NewCategoryHandler(s, services.Category),
	}
}
