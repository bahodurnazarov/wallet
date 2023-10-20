package handler

import (
	"github.com/bahodurnazarov/wallet/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) CheckWallet(c *gin.Context) {
	wallet := c.Query("wallet")
	if wallet == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Enter wallet number"})
		return
	}
	user, err := h.Service.CheckWallet(wallet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": user})
}

func (h *Handler) Replenishment(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	newUser, err := h.Service.Replenishment(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": newUser})
}
func (h *Handler) GetTransactions(c *gin.Context) {
	wallet := c.Query("wallet")
	if wallet == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Enter wallet number"})
		return
	}
	replenishment, err := h.Service.GetTransactions(wallet)
	log.Println("1111111111", replenishment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "Balance": replenishment})
}
func (h *Handler) Balance(c *gin.Context) {
	wallet := c.Query("wallet")
	if wallet == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Enter wallet number"})
		return
	}
	user, err := h.Service.Balance(wallet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "Balance": user.Balance})
}
