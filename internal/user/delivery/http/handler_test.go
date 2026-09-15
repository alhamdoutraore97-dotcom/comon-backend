package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alhamdoutraore97-dotcom/backend/internal/user/domain"
	mock_repo "github.com/alhamdoutraore97-dotcom/backend/internal/user/repository/mocks"
	"github.com/alhamdoutraore97-dotcom/backend/internal/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// setupRouter construit le routeur pour les tests
// setupRouter собирает маршрутизатор для тестов
func setupRouter(mockRepo *mock_repo.MockUserRepository) *gin.Engine {
	gin.SetMode(gin.TestMode)
	uc := usecase.NewUserUsecase(mockRepo)
	handler := NewUserHandler(uc)

	r := gin.New()
	r.GET("/users/:id", handler.GetUser)
	r.POST("/users", handler.CreateUser)
	return r
}

// ============ GET /users/:id ============

func TestHandler_GetUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().GetByID(gomock.Any(), "1").Return(&domain.User{ID: 1, Name: "Alice"}, nil)

	router := setupRouter(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Alice", resp["name"])
}

func TestHandler_GetUser_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	router := setupRouter(mockRepo)

	req := httptest.NewRequest(http.MethodGet, "/users/abc", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_GetUser_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().GetByID(gomock.Any(), "999").Return(nil, errors.New("not found"))

	router := setupRouter(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/users/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// ============ POST /users ============

func TestHandler_CreateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, u *domain.User) error {
		u.ID = 1
		return nil
	})

	router := setupRouter(mockRepo)

	body, _ := json.Marshal(map[string]string{
		"name":  "Bob",
		"email": "bob@test.com",
	})
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestHandler_CreateUser_BadJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	router := setupRouter(mockRepo)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString("{invalid"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_CreateUser_MissingFields(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	router := setupRouter(mockRepo)

	body, _ := json.Marshal(map[string]string{"name": "Bob"}) // email manquant
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
