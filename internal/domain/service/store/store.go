package store

import (
	"context"
	"fmt"
)




type Store struct {
	Ctx *contextController

	Map map[string] * contextController
}


func New(ctx context.Context) *Store {
	return &Store{
		Ctx: new_ctx(ctx),
		Map: make(map[string] * contextController),
	}
}


func (s *Store) StockExists(stock string) bool {
	_, ok := s.Map[stock]
	return ok
}


func (s *Store) StoreStock(stock string) error {
	if s.StockExists(stock) {
		return fmt.Errorf("stock %s already exists", stock)
	}

	s.Map[stock] = new_ctx(s.Ctx.ctx)
	return nil
}


func (s *Store) UnstoreStock(stock string) error {
	if !s.StockExists(stock) {
		return fmt.Errorf("stock %s not exists", stock)
	}
	s.Map[stock].cancel()
	delete(s.Map, stock)
	return nil
}


func (s *Store) Close() {
	s.Map = nil
	s.Ctx.cancel()
}