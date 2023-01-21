package database

import (
	"context"
	"fmt"
	"sync"

	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/domain"
)

type UserDB struct {
	// store - db in memory, key - userID, date
	store map[int64]domain.User
	mutex sync.RWMutex
}

func NewUserDB() (*UserDB, error) {
	return &UserDB{
		store: make(map[int64]domain.User),
	}, nil
}

func (db *UserDB) UserExist(ctx context.Context, userID int64) bool {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	_, ok := db.store[userID]

	return ok
}

func (db *UserDB) ChangeDefaultCurrency(ctx context.Context, userID int64, currency string) error {
	// Для того чтобы не было случаев, когда пользователь меняет выбранную валюту
	// и сразу же создает новую транзакцию с расходом, которая может создаться не
	// под текущей валютой из-за гонки обработки сообщений
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.store[userID] = domain.User{UserID: userID, DefaultCurrency: currency}

	return nil
}

func (db *UserDB) GetDefaultCurrency(ctx context.Context, userID int64) (string, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	if user, ok := db.store[userID]; ok {
		return user.DefaultCurrency, nil
	}

	return "", fmt.Errorf("user #%d not found", userID)
}
