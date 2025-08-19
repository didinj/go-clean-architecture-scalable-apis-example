package usecase_test

import (
	"errors"
	"testing"

	"github.com/didinj/go-clean-architecture/internal/entity"
	"github.com/didinj/go-clean-architecture/internal/usecase"
	"github.com/stretchr/testify/assert"
)

// --- Mock Repository ---
type mockUserRepository struct {
	users map[int64]*entity.User
	err   error
}

func (m *mockUserRepository) Create(user *entity.User) error {
	if m.err != nil {
		return m.err
	}
	user.ID = int64(len(m.users) + 1)
	m.users[user.ID] = user
	return nil
}

func (m *mockUserRepository) GetByID(id int64) (*entity.User, error) {
	if user, ok := m.users[id]; ok {
		return user, nil
	}
	return nil, errors.New("not found")
}

func (m *mockUserRepository) GetAll() ([]*entity.User, error) {
	var allUsers []*entity.User
	for _, user := range m.users {
		allUsers = append(allUsers, user)
	}
	return allUsers, nil
}

// --- Tests ---
func TestRegisterUser_Success(t *testing.T) {
	repo := &mockUserRepository{users: make(map[int64]*entity.User)}
	uc := usecase.NewUserUsecase(repo)

	user := &entity.User{Name: "Alice", Email: "alice@example.com"}
	err := uc.CreateUser(user)

	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)
}

func TestRegisterUser_Failure(t *testing.T) {
	repo := &mockUserRepository{err: errors.New("db error"), users: make(map[int64]*entity.User)}
	uc := usecase.NewUserUsecase(repo)

	user := &entity.User{Name: "Bob", Email: "bob@example.com"}
	err := uc.CreateUser(user)

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
}
