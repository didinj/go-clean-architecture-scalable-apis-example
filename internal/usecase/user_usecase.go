package usecase

import (
	"github.com/didinj/go-clean-architecture/internal/entity"
	"github.com/didinj/go-clean-architecture/internal/repository"
)

// UserUsecase contains the business logic for users.
type UserUsecase struct {
	repo repository.UserRepository
}

// NewUserUsecase creates a new UserUsecase instance.
func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// CreateUser handles the business logic of creating a user.
func (uc *UserUsecase) CreateUser(user *entity.User) error {
	if user.Email == "" {
		return ErrInvalidEmail
	}
	return uc.repo.Create(user)
}

// GetUserByID retrieves a user by ID.
func (uc *UserUsecase) GetUserByID(id int64) (*entity.User, error) {
	return uc.repo.GetByID(id)
}

// ListUsers retrieves all users.
func (uc *UserUsecase) ListUsers() ([]*entity.User, error) {
	return uc.repo.GetAll()
}
