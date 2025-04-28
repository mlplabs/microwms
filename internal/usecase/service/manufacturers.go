package service

import (
	"context"
	"github.com/mlplabs/mwms-core/whs/model"
)

func (s *WhsService) GetManufacturerById(ctx context.Context, userId int64) (*model.Manufacturer, error) {
	return s.manufacturersCatalog.GetById(ctx, userId)
}
func (s *WhsService) GetManufacturers(ctx context.Context, offset int, limit int, search string) ([]model.Manufacturer, int64, error) {
	return s.manufacturersCatalog.GetItems(ctx, offset, limit, search)
}
func (s *WhsService) CreateManufacturer(ctx context.Context, mnf *model.Manufacturer) (int64, error) {
	return s.manufacturersCatalog.Create(ctx, mnf)
}
func (s *WhsService) UpdateManufacturer(ctx context.Context, mnf *model.Manufacturer) (int64, error) {
	return s.manufacturersCatalog.Update(ctx, mnf)
}
func (s *WhsService) DeleteManufacturer(ctx context.Context, mnfId int64) error {
	_, err := s.manufacturersCatalog.GetById(ctx, mnfId)
	if err != nil {
		return err
	}
	return s.manufacturersCatalog.Delete(ctx, mnfId)
}
func (s *WhsService) GetManufacturerSuggestion(ctx context.Context, text string, limit int) ([]model.Suggestion, error) {
	return s.manufacturersCatalog.Suggest(ctx, text, limit)
}

func (s *WhsService) FindOrCreateManufacturer(ctx context.Context, manufacturerName string) (*model.Manufacturer, error) {
	mnfItems, err := s.manufacturersCatalog.FindByName(ctx, manufacturerName)
	if err != nil {
		return nil, err
	}
	if len(mnfItems) > 0 {
		return &mnfItems[0], nil
	} else {
		mnfId := int64(0)
		mnfId, err = s.manufacturersCatalog.Create(ctx, &model.Manufacturer{
			Name: manufacturerName,
		})
		if err != nil {
			return nil, err
		}
		return &model.Manufacturer{
			Id:   mnfId,
			Name: manufacturerName,
		}, nil
	}
}
