package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/testing", h.testHanlder)
}

func (h *Handler) testHanlder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "test complete",
	})
}
