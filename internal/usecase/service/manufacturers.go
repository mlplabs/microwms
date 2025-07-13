package service

import (
	"context"
	"github.com/mlplabs/mwms-core/whs/model"
)

func (s *WhsService) GetManufacturerById(ctx context.Context, userId int64) (*model.Manufacturer, error) {
	return s.storage.GetManufacturerById(ctx, userId)
}
func (s *WhsService) GetManufacturers(ctx context.Context, offset int, limit int, search string) ([]model.Manufacturer, int64, error) {
	return s.storage.GetManufacturersItems(ctx, offset, limit, search)
}
func (s *WhsService) CreateManufacturer(ctx context.Context, mnf *model.Manufacturer) (int64, error) {
	return s.storage.CreateManufacturer(ctx, mnf)
}
func (s *WhsService) UpdateManufacturer(ctx context.Context, mnf *model.Manufacturer) (int64, error) {
	return s.storage.UpdateManufacturer(ctx, mnf)
}
func (s *WhsService) DeleteManufacturer(ctx context.Context, mnfId int64) error {
	_, err := s.storage.GetManufacturerById(ctx, mnfId)
	if err != nil {
		return err
	}
	return s.storage.DeleteManufacturer(ctx, mnfId)
}
func (s *WhsService) GetManufacturerSuggestion(ctx context.Context, text string, limit int) ([]model.Suggestion, error) {
	return s.storage.ManufacturersSuggest(ctx, text, limit)
}

func (s *WhsService) FindOrCreateManufacturer(ctx context.Context, manufacturerName string) (*model.Manufacturer, error) {
	mnfItems, err := s.storage.FindManufacturersByName(ctx, manufacturerName)
	if err != nil {
		return nil, err
	}
	if len(mnfItems) > 0 {
		return &mnfItems[0], nil
	} else {
		mnfId := int64(0)
		mnfId, err = s.storage.CreateManufacturer(ctx, &model.Manufacturer{
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
