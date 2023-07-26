package persistence

import (
	"context"

	"github.com/Goboolean/fetch-server/internal/domain/entity"
	"github.com/Goboolean/fetch-server/internal/domain/port"
	"github.com/Goboolean/fetch-server/internal/domain/port/out"
)




type MockAdapter struct {
	db map[string]int
}

func NewMockAdapter() out.StockPersistencePort {
	return &MockAdapter{db: make(map[string]int)}
}


func (a *MockAdapter) StoreStock(tx port.Transactioner, stockId string, agg *entity.StockAggregate) error {
	if _, ok := a.db[stockId]; !ok {
		a.db[stockId] = 0
	}
	a.db[stockId]++
	return nil
}

func (a *MockAdapter) StoreStockBatch(tx port.Transactioner, stockId string, batch []*entity.StockAggregate) error {
	if _, ok := a.db[stockId]; !ok {
		a.db[stockId] = 0
	}
	a.db[stockId] += len(batch)
	return nil
}

func (a *MockAdapter) CreateStoringStartedLog(context.Context, string) error {
	return nil
}

func (a *MockAdapter) CreateStoringFailedLog(context.Context, string) error {
	return nil
}

func (a *MockAdapter) CreateStoringStoppedLog(context.Context, string) error {
	return nil
}


func (a *MockAdapter) GetStoredStockCount(stockId string) int {
	if _, ok := a.db[stockId]; !ok {
		a.db[stockId] = 0
	}
	return a.db[stockId]
}


func (a *MockAdapter) Clear() {
	for k := range a.db {
		delete(a.db, k)
	}
}