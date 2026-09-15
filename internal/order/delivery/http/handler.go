package http

import (
	"net/http"
	"strconv"

	"github.com/alhamdoutraore97-dotcom/backend/internal/order/usecase"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	usecase *usecase.OrderUsecase
}

func NewOrderHandler(usecase *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: usecase}
}

// GetOrder godoc
// @Summary Get order by ID / Получить заказ по ID
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} domain.Order
// @Failure 404 {object} map[string]string
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	order, err := h.usecase.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// CreateOrder godoc
// @Summary Create order / Создать заказ
// @Accept json
// @Produce json
// @Param order body domain.Order true "Order data"
// @Success 201 {object} domain.Order
// @Failure 400 {object} map[string]string
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req struct {
		UserID   int    `json:"user_id"`
		Product  string `json:"product"`
		Quantity int    `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.usecase.CreateOrder(req.UserID, req.Product, req.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}
