package services

import (
	"context"

	"github.com/alielmi98/golang-expense-tracker-api/api/dto"
	"github.com/alielmi98/golang-expense-tracker-api/common"

	"github.com/alielmi98/golang-expense-tracker-api/data/models"
	"github.com/alielmi98/golang-expense-tracker-api/repository"
)

type ExpenseTrackerService interface {
	CreateExpense(ctx context.Context, expense *dto.CreateExpenseRequest) (*dto.ExpenseResponse, error)
}

type expenseTrackerService struct {
	repo repository.ExpenseTrackerRepository
}

func NewExpenseTrackerService() ExpenseTrackerService {
	return &expenseTrackerService{
		repo: repository.NewExpenseTrackerRepository(),
	}
}

func (s *expenseTrackerService) CreateExpense(ctx context.Context, expense *dto.CreateExpenseRequest) (*dto.ExpenseResponse, error) {
	userId := uint(1)
	model, err := common.TypeConverter[models.Expense](expense)
	if err != nil {
		return &dto.ExpenseResponse{}, err
	}

	model.UserID = userId

	err = s.repo.CreateExpense(ctx, model)
	if err != nil {
		return &dto.ExpenseResponse{}, err
	}

	response, err := common.TypeConverter[dto.ExpenseResponse](model)
	if err != nil {
		return &dto.ExpenseResponse{}, err
	}

	return response, err
}
