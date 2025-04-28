package service

import (
	"context"
	"github.com/mlplabs/mwms-core/whs/model"
)

func (s *WhsService) GetWarehouseById(ctx context.Context, whsId int64) (*model.Warehouse, error) {
	return s.warehousesCatalog.GetById(ctx, whsId)
}

func (s *WhsService) GetWarehouses(ctx context.Context, offset int, limit int) ([]model.Warehouse, int64, error) {
	return s.warehousesCatalog.GetItems(ctx, offset, limit)
}

func (s *WhsService) CreateWarehouse(ctx context.Context, whs *model.Warehouse) (int64, error) {
	return s.warehousesCatalog.Create(ctx, whs)
}

func (s *WhsService) UpdateWarehouse(ctx context.Context, whs *model.Warehouse) (int64, error) {
	return s.warehousesCatalog.Update(ctx, whs)
}

func (s *WhsService) DeleteWarehouse(ctx context.Context, whsId int64) error {
	_, err := s.warehousesCatalog.GetById(ctx, whsId)
	if err != nil {
		return err
	}
	return s.warehousesCatalog.Delete(ctx, whsId)
}

func (s *WhsService) GetWarehouseSuggestion(ctx context.Context, text string, limit int) ([]model.Suggestion, error) {
	return s.warehousesCatalog.Suggest(ctx, text, limit)
}
