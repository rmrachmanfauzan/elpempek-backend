package service

import (
	"context"

	"github.com/rmrachmanfauzan/elpempek/internal/dto"
	repo "github.com/rmrachmanfauzan/elpempek/internal/repository"
	"github.com/rmrachmanfauzan/elpempek/internal/model"
)

type UserService interface {
	CreateUser(ctx context.Context ,req dto.CreateUserRequest) (*model.User, error)
	GetUserByID(ctx context.Context ,id uint) (*model.User, error)
}

type userService struct {
	repo repo.UserRepository
}

func NewUserService(r repo.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*model.User, error) {
	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}
