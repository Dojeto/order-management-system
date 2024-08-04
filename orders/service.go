package main

import "context"

type Service struct {
	store OrderStore
}

func (serv Service) CreateOrder(ctx context.Context) error {
	return nil
}
