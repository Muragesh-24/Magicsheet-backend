package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service{
	return &Service{
		repo: repo,
	}
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error){
	
	user, err := s.repo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.HashPassword),
		[]byte(req.Password),
	)


	if err != nil {
		return nil, ErrInvalidCredentials
	}

	
	return &LoginResponse{
		AccessToken: "temp - token",
	}, nil 
}