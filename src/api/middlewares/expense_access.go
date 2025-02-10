package middlewares

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/alielmi98/golang-expense-tracker-api/api/helper"
	"github.com/alielmi98/golang-expense-tracker-api/constants"
	"github.com/alielmi98/golang-expense-tracker-api/services"
	"github.com/gin-gonic/gin"
)

func AuthorizeExpenseAccess() gin.HandlerFunc {
	service := services.NewExpenseTrackerService()
	return func(c *gin.Context) {
		userId := int(c.Value(constants.UserIdKey).(float64))
		expenseId, _ := strconv.Atoi(c.Params.ByName("id"))

		expense, err := service.GetExpenseByID(c.Request.Context(), expenseId)
		if err != nil {
			c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
				helper.GenerateBaseResponseWithError(nil, false, helper.NotFoundError, err))
			return
		}

		if int(expense.UserID) != userId {
			err := errors.New("access denied")
			c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponseWithError(
				nil, false, helper.AuthError, err,
			))
			return
		}

		c.Next()
	}
}
