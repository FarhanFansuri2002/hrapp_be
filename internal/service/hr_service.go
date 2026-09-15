package service

import (
	"context"

	"server/internal/model"
	"server/internal/repository"
)

type HRService struct {
	store *repository.MySQLStore
}

func NewHRService(store *repository.MySQLStore) *HRService {
	return &HRService{store: store}
}

func (service *HRService) ListEmployees(ctx context.Context) ([]model.Employee, error) {
	return service.store.ListEmployees(ctx)
}
