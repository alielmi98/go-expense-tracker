package services

import (
	"context"
	"time"

	"github.com/alielmi98/golang-expense-tracker-api/api/dto"
	"github.com/alielmi98/golang-expense-tracker-api/common"
	"github.com/alielmi98/golang-expense-tracker-api/constants"

	"github.com/alielmi98/golang-expense-tracker-api/data/models"
	"github.com/alielmi98/golang-expense-tracker-api/repository"
)

type ExpenseTrackerService interface {
	CreateExpense(ctx context.Context, expense *dto.CreateExpenseRequest) (*dto.ExpenseResponse, error)
	UpdateExpense(ctx context.Context, id int, expense *dto.UpdateExpenseRequest) (*dto.ExpenseResponse, error)
	DeleteExpense(ctx context.Context, id int) error
	GetExpenseByID(ctx context.Context, id int) (*dto.ExpenseResponse, error)
	ListExpenses(ctx context.Context, startDate, endDate time.Time) ([]dto.ExpenseResponse, error)
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
	userId := uint(ctx.Value(constants.UserIdKey).(float64))

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

func (s *expenseTrackerService) UpdateExpense(ctx context.Context, id int, expense *dto.UpdateExpenseRequest) (*dto.ExpenseResponse, error) {
	model, err := common.TypeConverter[models.Expense](expense)
	if err != nil {
		return &dto.ExpenseResponse{}, err
	}

	err = s.repo.UpdateExpense(ctx, id, model)
	if err != nil {
		return &dto.ExpenseResponse{}, err
	}
	updatedModel, _ := s.repo.GetExpenseByID(ctx, id)
	response, err := common.TypeConverter[dto.ExpenseResponse](updatedModel)
	if err != nil {
		return &dto.ExpenseResponse{}, err
	}
	return response, err

}

func (s *expenseTrackerService) DeleteExpense(ctx context.Context, id int) error {
	err := s.repo.DeleteExpense(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *expenseTrackerService) GetExpenseByID(ctx context.Context, id int) (*dto.ExpenseResponse, error) {
	model, err := s.repo.GetExpenseByID(ctx, id)
	if err != nil {
		return &dto.ExpenseResponse{}, err
	}
	response, err := common.TypeConverter[dto.ExpenseResponse](model)
	if err != nil {
		return &dto.ExpenseResponse{}, err
	}
	return response, nil
}

func (s *expenseTrackerService) ListExpenses(ctx context.Context, startDate, endDate time.Time) ([]dto.ExpenseResponse, error) {
	userId := uint(ctx.Value(constants.UserIdKey).(float64))
	models, err := s.repo.ListExpenses(ctx, userId, startDate, endDate)
	if err != nil {
		return nil, err
	}

	var responses []dto.ExpenseResponse
	for _, model := range models {
		response, err := common.TypeConverter[dto.ExpenseResponse](model)
		if err != nil {
			return nil, err
		}
		responses = append(responses, *response)
	}
	return responses, nil
}
