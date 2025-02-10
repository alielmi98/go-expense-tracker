package routers

import (
	"github.com/alielmi98/golang-expense-tracker-api/api/handlers"
	"github.com/alielmi98/golang-expense-tracker-api/config"
	"github.com/gin-gonic/gin"
)

func Expense(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewExpenseTrackerHandler()
	router.POST("/", h.CreateExpense)
	router.PUT("/:id", h.UpdateExpense)
	router.DELETE("/:id", h.DeleteExpense)
	router.GET("/:id", h.GetExpenseByID)
	router.GET("/", h.ListExpenses)

}
