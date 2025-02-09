package handlers

import (
	"net/http"

	"github.com/alielmi98/golang-expense-tracker-api/api/dto"
	"github.com/alielmi98/golang-expense-tracker-api/api/helper"
	"github.com/alielmi98/golang-expense-tracker-api/services"
	"github.com/gin-gonic/gin"
)

type expenseTrackerHandler struct {
	expenseService services.ExpenseTrackerService
}

func NewExpenseTrackerHandler() *expenseTrackerHandler {
	return &expenseTrackerHandler{
		expenseService: services.NewExpenseTrackerService(),
	}
}

// CreateTodo godoc
// @Summary Create a Expense
// @Description Create a new Expense
// @Tags Expenses
// @Accept json
// @produces json
// @Param Request body dto.CreateExpenseRequest true "Create a Expense"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.ExpenseResponse} "Expense response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/expense/ [post]
// @Security AuthBearer
func (h *expenseTrackerHandler) CreateExpense(c *gin.Context) {

	var createReqDTO dto.CreateExpenseRequest
	if err := c.ShouldBindJSON(&createReqDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	response, err := h.expenseService.CreateExpense(c, &createReqDTO)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}
