package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alhamdoutraore97-dotcom/backend/internal/order/domain"
	mock_repo "github.com/alhamdoutraore97-dotcom/backend/internal/order/repository/mocks"
	"github.com/alhamdoutraore97-dotcom/backend/internal/order/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setupRouter(mockRepo *mock_repo.MockOrderRepository) *gin.Engine {
	gin.SetMode(gin.TestMode)
	uc := usecase.NewOrderUsecase(mockRepo)
	handler := NewOrderHandler(uc)
	r := gin.New()
	r.GET("/orders/:id", handler.GetOrder)
	r.POST("/orders", handler.CreateOrder)
	return r
}

func TestHandler_GetOrder_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	mockRepo.EXPECT().GetByID(gomock.Any(), "1").Return(&domain.Order{ID: 1, Product: "Book"}, nil)

	router := setupRouter(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/orders/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Book", resp["product"])
}

func TestHandler_GetOrder_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	router := setupRouter(mockRepo)

	req := httptest.NewRequest(http.MethodGet, "/orders/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_GetOrder_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	mockRepo.EXPECT().GetByID(gomock.Any(), "99").Return(nil, errors.New("not found"))
	router := setupRouter(mockRepo)

	req := httptest.NewRequest(http.MethodGet, "/orders/99", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestHandler_CreateOrder_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
	router := setupRouter(mockRepo)

	body, _ := json.Marshal(map[string]interface{}{
		"user_id": 1, "product": "Book", "quantity": 2,
	})
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestHandler_CreateOrder_BadJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	router := setupRouter(mockRepo)

	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString("{bad"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_CreateOrder_InvalidData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	router := setupRouter(mockRepo)

	body, _ := json.Marshal(map[string]interface{}{
		"user_id": 0, "product": "", "quantity": 0,
	})
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
