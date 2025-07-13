package service

import (
	"context"
	"github.com/mlplabs/mwms-core/whs/model"
)

func (s *WhsService) GetWarehouseById(ctx context.Context, whsId int64) (*model.Warehouse, error) {
	return s.storage.GetWarehouseById(ctx, whsId)
}

func (s *WhsService) GetWarehouses(ctx context.Context, offset int, limit int) ([]model.Warehouse, int64, error) {
	return s.storage.GetWarehousesItems(ctx, offset, limit)
}

func (s *WhsService) CreateWarehouse(ctx context.Context, whs *model.Warehouse) (int64, error) {
	return s.storage.CreateWarehouse(ctx, whs)
}

func (s *WhsService) UpdateWarehouse(ctx context.Context, whs *model.Warehouse) (int64, error) {
	return s.storage.UpdateWarehouse(ctx, whs)
}

func (s *WhsService) DeleteWarehouse(ctx context.Context, whsId int64) error {
	_, err := s.storage.GetWarehouseById(ctx, whsId)
	if err != nil {
		return err
	}
	return s.storage.DeleteWarehouse(ctx, whsId)
}

func (s *WhsService) GetWarehouseSuggestion(ctx context.Context, text string, limit int) ([]model.Suggestion, error) {
	return s.storage.WarehousesSuggest(ctx, text, limit)
}
