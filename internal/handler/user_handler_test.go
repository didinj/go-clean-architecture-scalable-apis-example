package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/didinj/go-clean-architecture/internal/entity"
	"github.com/didinj/go-clean-architecture/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Compile-time check: ensure mockUserUsecase implements usecase.UserUsecase
var _ usecase.UserUsecase = (*mockUserUsecase)(nil)

// Mock Usecase
type mockUserUsecase struct{}

func (m *mockUserUsecase) RegisterUser(user *entity.User) error {
	user.ID = 99
	return nil
}
func (m *mockUserUsecase) GetUser(id int64) (*entity.User, error) {
	return &entity.User{ID: id, Name: "Test", Email: "test@example.com"}, nil
}

func TestRegisterUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	h := NewUserHandler(&mockUserUsecase{})
	h.RegisterRoutes(r)

	body := `{"name":"Charlie","email":"charlie@example.com"}`
	req, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"id":99`)
}

func TestGetUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	h := NewUserHandler(&mockUserUsecase{})
	h.RegisterRoutes(r)

	// Create a request for GET /users/42
	req, _ := http.NewRequest(http.MethodGet, "/users/42", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"id":42`)
	assert.Contains(t, w.Body.String(), `"name":"Test"`)
	assert.Contains(t, w.Body.String(), `"email":"test@example.com"`)
}
