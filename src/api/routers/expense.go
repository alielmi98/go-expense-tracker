package routers

import (
	"github.com/alielmi98/golang-expense-tracker-api/api/handlers"
	"github.com/alielmi98/golang-expense-tracker-api/api/middlewares"
	"github.com/alielmi98/golang-expense-tracker-api/config"

	"github.com/gin-gonic/gin"
)

func Expense(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewExpenseTrackerHandler()
	router.POST("/", h.CreateExpense)
	router.PUT("/:id", middlewares.AuthorizeExpenseAccess(), h.UpdateExpense)
	router.DELETE("/:id", middlewares.AuthorizeExpenseAccess(), h.DeleteExpense)
	router.GET("/:id", middlewares.AuthorizeExpenseAccess(), h.GetExpenseByID)
	router.GET("/", h.ListExpenses)

}
