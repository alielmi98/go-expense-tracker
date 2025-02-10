package repository

import (
	"context"
	"log"

	"github.com/alielmi98/golang-expense-tracker-api/constants"
	"github.com/alielmi98/golang-expense-tracker-api/data/db"
	"github.com/alielmi98/golang-expense-tracker-api/data/models"

	"gorm.io/gorm"
)

type ExpenseTrackerRepository interface {
	CreateExpense(ctx context.Context, model *models.Expense) error
	UpdateExpense(ctx context.Context, id int, model *models.Expense) error
	DeleteExpense(ctx context.Context, id int) error
	GetExpenseByID(ctx context.Context, id int) (*models.Expense, error)
}

type expenseTrackerRepository struct {
	db *gorm.DB
}

func NewExpenseTrackerRepository() ExpenseTrackerRepository {
	return &expenseTrackerRepository{
		db: db.GetDb(),
	}
}

func (r *expenseTrackerRepository) CreateExpense(ctx context.Context, model *models.Expense) error {
	tx := r.db.WithContext(ctx).Begin()
	err := tx.
		Create(model).
		Error
	if err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Postgres, constants.Insert, err.Error())
		return err
	}
	tx.Commit()

	return nil
}

func (r *expenseTrackerRepository) UpdateExpense(ctx context.Context, id int, model *models.Expense) error {
	tx := r.db.WithContext(ctx).Begin()
	if err := tx.Model(model).
		Where("id = ? ", id).
		Updates(model).
		Error; err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Postgres, constants.Update, err.Error())
		return err
	}
	tx.Commit()

	return nil
}

func (r *expenseTrackerRepository) DeleteExpense(ctx context.Context, id int) error {
	tx := r.db.WithContext(ctx).Begin()

	model := new(models.Expense)
	if err := tx.Where("id = ?", id).Delete(model).Error; err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Postgres, constants.Delete, err.Error())
		return err
	}
	tx.Commit()

	return nil
}

func (r *expenseTrackerRepository) GetExpenseByID(ctx context.Context, id int) (*models.Expense, error) {
	var expense models.Expense
	if err := r.db.WithContext(ctx).First(&expense, id).Error; err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Postgres, constants.Select, err.Error())
		return nil, err
	}
	return &expense, nil
}
