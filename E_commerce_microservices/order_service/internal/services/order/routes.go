package order

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"trann/ecom/order_service/internal/models"
	"trann/ecom/order_service/internal/types"
	"trann/ecom/order_service/internal/utils"
)

type Handler struct {
	orderStore types.OrderRepository
}

func NewHandler(orderStore types.OrderRepository) *Handler {
	return &Handler{
		orderStore: orderStore,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/api/v1/orders", h.hanldeGetOrders)
	router.POST("/api/v1/orders", h.handlePostCreateOrder)
	router.GET("/api/v1/orders/:id", h.handleGetOrderById)
	router.POST("/api/v1/orders/:id/order-details", h.handlePostOrderDetailsToOrderById)
}

func (h *Handler) hanldeGetOrders(c *gin.Context) {
	ctx := c.Request.Context()

	// Get orders from the repository
	orders, err := h.orderStore.GetOrders(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert data models to payload models
	var orderPayloads []types.OrderPayload
	for _, order := range orders {
		orderPayloads = append(orderPayloads, utils.ToOrderPayload(order))
	}
	// Respond with the payload
	c.JSON(http.StatusOK, orderPayloads)
}

func (h *Handler) handlePostCreateOrder(c *gin.Context) {
	var createOrderPayload types.CreateOrderPayload

	// Bind JSON payload to the struct
	if err := c.ShouldBindJSON(&createOrderPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert payload to data model
	order := models.Order{
		Status:        "In Cart",
		LastUpdatedAt: time.Now(),
		CreatedAt:     time.Now(),
	}
	for _, detail := range createOrderPayload.OrderDetails {
		order.OrderDetails = append(order.OrderDetails, models.OrderDetail{
			ItemID:   detail.ItemID,
			Quantity: detail.Quantity,
		})
	}

	// Create order in the repository
	if err := h.orderStore.CreateOrder(c.Request.Context(), &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created order ID
	c.JSON(
		http.StatusCreated,
		gin.H{"msg": fmt.Sprintf("new order created with id : %d", order.ID)},
	)
}

func (h *Handler) handleGetOrderById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	order, err := h.orderStore.GetOrderById(c.Request.Context(), idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var orderPayload types.OrderPayload
	orderPayload = utils.ToOrderPayload(order)

	c.JSON(http.StatusOK, orderPayload)
}

func (h *Handler) handlePostOrderDetailsToOrderById(c *gin.Context) {
	var req types.AddOrderDetailsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	orderId := c.Param("id")
	idInt, err := strconv.Atoi(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	order, err := h.orderStore.GetOrderById(c.Request.Context(), idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var orderDetails []*models.OrderDetail
	for _, v := range req.OrderDetails {
		orderDetail := models.OrderDetail{
			OrderID:       order.ID,
			ItemID:        v.ItemID,
			Quantity:      v.Quantity,
			CreatedAt:     time.Now(),
			LastUpdatedAt: time.Now(),
		}
		orderDetails = append(orderDetails, &orderDetail)
	}
	if err := h.orderStore.CreateOrderDetailsToOrder(c.Request.Context(), idInt, orderDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	err = h.orderStore.UpdateOrderTime(c.Request.Context(), &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, utils.ToOrderPayload(order))
}
