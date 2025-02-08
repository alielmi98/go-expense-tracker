package routers

import (
	"github.com/alielmi98/golang-expense-tracker-api/api/handlers"
	"github.com/alielmi98/golang-expense-tracker-api/config"

	"github.com/gin-gonic/gin"
)

func Token(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewTokensHandler(cfg)
	router.POST("/refresh-token", h.RefreshToken)
}
