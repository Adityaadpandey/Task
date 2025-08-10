package service

import (
	"github.com/adityaadpandey/go-boilerplate/internal/lib/job"
	"github.com/adityaadpandey/go-boilerplate/internal/repository"
	"github.com/adityaadpandey/go-boilerplate/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
