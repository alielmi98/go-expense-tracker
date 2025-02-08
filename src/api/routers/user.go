package routers

import (
	"github.com/alielmi98/golang-expense-tracker-api/api/handlers"
	"github.com/alielmi98/golang-expense-tracker-api/config"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)
	router.POST("/login-by-username", h.LoginByUsername)
	router.POST("/register-by-username", h.RegisterByUsername)

}
