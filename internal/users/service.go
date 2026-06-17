package users

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, req CreateUserRequest) (UserResponse, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Role = strings.TrimSpace(req.Role)

	if req.Name == "" {
		return UserResponse{}, errors.New("name is required")
	}

	if req.Email == "" {
		return UserResponse{}, errors.New("email is required")
	}

	if req.Password == "" {
		return UserResponse{}, errors.New("password is required")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserResponse{}, err
	}

	user := NewUser(
		req.Name,
		req.Email,
		string(passwordHash),
		req.Role,
	)

	createdUser, err := s.repository.Create(ctx, user)
	if err != nil {
		return UserResponse{}, err
	}

	return toUserResponse(createdUser), nil
}

func (s *Service) FindAll(ctx context.Context) ([]UserResponse, error) {
	users, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]UserResponse, 0, len(users))

	for _, user := range users {
		response = append(response, toUserResponse(user))
	}

	return response, nil
}

func toUserResponse(user User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}
