package handler

import (
	"github.com/bahodurnazarov/wallet/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	if err := h.Service.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Created New User"})
}
