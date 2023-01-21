package database

import (
	"context"
	"time"

	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/domain"
)

type ExpenseDB struct {
	// store - db in memory, key - userID, date
	store map[int64][]domain.Expense
}

func NewExpenseDB() (*ExpenseDB, error) {
	return &ExpenseDB{
		store: make(map[int64][]domain.Expense),
	}, nil
}

// AddExpense - добавление нового рассхода
func (db *ExpenseDB) AddExpense(ctx context.Context, userID int64, kopecks int64, title string, date time.Time) error {
	db.store[userID] = append(db.store[userID], domain.Expense{
		Title:  title,
		Date:   date,
		Amount: kopecks,
	})
	return nil
}

// GetExpenses - получение всех рассходов пользователя
func (db *ExpenseDB) GetExpenses(ctx context.Context, userID int64) ([]domain.Expense, error) {
	return db.store[userID], nil
}
