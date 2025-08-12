package v1

import (
	"github.com/adityaadpandey/go-boilerplate/internal/handler"
	"github.com/adityaadpandey/go-boilerplate/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterV1Routes(router *echo.Group, handlers *handler.Handlers, middleware *middleware.Middlewares) {

	// Register Todo routes
	registerTodoRoutes(router, handlers.TodoHandler, handlers.CommentHandler, middleware.Auth)

	// Register Comment routes
	registerCommentRoutes(router, handlers.CommentHandler, middleware.Auth)

	// Register Category routes
	registerCategoryRoutes(router, handlers.CategoryHandler, middleware.Auth)
}
