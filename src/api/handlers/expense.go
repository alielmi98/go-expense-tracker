package handlers

import (
	"net/http"
	"strconv"
	"time"

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

// UpdateExpense godoc
// @Summary Update a Expense
// @Description Update a Expense by Id
// @Tags Expenses
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateExpenseRequest true "Update a Expense"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.ExpenseResponse} "Expense response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/expense/{id} [put]
// @Security AuthBearer
func (h *expenseTrackerHandler) UpdateExpense(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var updateReqDTO dto.UpdateExpenseRequest
	if err := c.ShouldBindJSON(&updateReqDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	response, err := h.expenseService.UpdateExpense(c, id, &updateReqDTO)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}

// DeleteExpense godoc
// @Summary Delete a Expense
// @Description Delete a Expense by Id
// @Tags Expenses
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 204 {object} helper.BaseHttpResponse "No content"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/expense/{id} [delete]
// @Security AuthBearer
func (h *expenseTrackerHandler) DeleteExpense(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	err := h.expenseService.DeleteExpense(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusNoContent, helper.GenerateBaseResponse(nil, true, 0))
}

// GetExpenseByID godoc
// @Summary Get a Expense
// @Description Get a Expense by Id
// @Tags Expenses
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ExpenseResponse} "Expense response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/expense/{id} [get]
// @Security AuthBearer
func (h *expenseTrackerHandler) GetExpenseByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	response, err := h.expenseService.GetExpenseByID(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// ListExpenses godoc
// @Summary List Expenses
// @Description List and filter expenses
// @Tags Expenses
// @Accept json
// @produces json
// @Param filter query string false "Filter (past_week, past_month, last_3_months)"
// @Param startDate query string false "Start Date (YYYY-MM-DD)"
// @Param endDate query string false "End Date (YYYY-MM-DD)"
// @Success 200 {object} helper.BaseHttpResponse{result=[]dto.ExpenseResponse} "Expense response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/expense/ [get]
// @Security AuthBearer
func (h *expenseTrackerHandler) ListExpenses(c *gin.Context) {
	filter := c.Query("filter")
	startDateStr := c.Query("startDate")
	endDateStr := c.Query("endDate")

	var startDate, endDate time.Time
	var err error

	switch filter {
	case "past_week":
		startDate = time.Now().AddDate(0, 0, -7)
		endDate = time.Now()
	case "past_month":
		startDate = time.Now().AddDate(0, -1, 0)
		endDate = time.Now()
	case "last_3_months":
		startDate = time.Now().AddDate(0, -3, 0)
		endDate = time.Now()
	default:
		if startDateStr != "" {
			startDate, err = time.Parse("2006-01-02", startDateStr)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest,
					helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
				return
			}
		} else {
			startDate = time.Now().AddDate(0, 0, -7) // Default to past week
		}

		if endDateStr != "" {
			endDate, err = time.Parse("2006-01-02", endDateStr)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest,
					helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
				return
			}
		} else {
			endDate = time.Now()
		}
	}

	response, err := h.expenseService.ListExpenses(c, startDate, endDate)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
