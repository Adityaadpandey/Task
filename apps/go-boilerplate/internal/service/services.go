package service

import (
	"fmt"

	"github.com/adityaadpandey/go-boilerplate/internal/lib/aws"
	"github.com/adityaadpandey/go-boilerplate/internal/lib/job"
	"github.com/adityaadpandey/go-boilerplate/internal/repository"
	"github.com/adityaadpandey/go-boilerplate/internal/server"
)

type Services struct {
	Auth     *AuthService
	Job      *job.JobService
	Todo     *TodoService
	Comment  *CommentService
	Category *CategoryService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)
	s.Job.SetAuthService(authService)
	awsClient, err := aws.NewAWS(s)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS client: %w", err)
	}
	return &Services{
		Job:      s.Job,
		Auth:     authService,
		Category: NewCategoryService(s, repos.Category),
		Comment:  NewCommentService(s, repos.Comment, repos.Todo),
		Todo:     NewTodoService(s, repos.Todo, repos.Category, awsClient),
	}, nil
}
