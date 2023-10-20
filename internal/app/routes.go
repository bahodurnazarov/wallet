package app

import (
	"github.com/bahodurnazarov/wallet/internal/handler"
	"github.com/bahodurnazarov/wallet/pkg/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(cors.CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	r.POST("/createUser", h.CreateUser)
	r.POST("/checkWallet", h.CheckWallet)
	r.POST("/replenishment", h.Replenishment)
	r.POST("/getTransactions", h.GetTransactions)
	r.POST("/balance", h.Balance)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    "works!",
			"message": "route not found!"})
	})

	return r
}
