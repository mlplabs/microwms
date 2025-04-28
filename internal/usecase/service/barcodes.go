package service

import (
	"context"
	"github.com/mlplabs/mwms-core/whs/model"
)

func (s *WhsService) GetBarcodes(ctx context.Context, offset int, limit int) ([]model.Barcode, int64, error) {
	return s.barcodesCatalog.GetItems(ctx, offset, limit)
}

func (s *WhsService) GetBarcodesByOwner(ctx context.Context, offset int, limit int, ownerId int64, ownerRef string) ([]model.Barcode, int64, error) {
	return s.barcodesCatalog.GetItemsByOwner(ctx, offset, limit, ownerId, ownerRef)
}

func (s *WhsService) GetBarcodeById(ctx context.Context, itemId int64) (*model.Barcode, error) {
	return s.barcodesCatalog.GetById(ctx, itemId)
}

func (s *WhsService) CreateBarcode(ctx context.Context, bc *model.Barcode) (int64, error) {
	return s.barcodesCatalog.Create(ctx, bc)
}

func (s *WhsService) UpdateBarcode(ctx context.Context, bc *model.Barcode) (int64, error) {
	return s.barcodesCatalog.Update(ctx, bc)
}

func (s *WhsService) DeleteBarcode(ctx context.Context, itemId int64) error {
	_, err := s.barcodesCatalog.GetById(ctx, itemId)
	if err != nil {
		return err
	}
	return s.barcodesCatalog.Delete(ctx, itemId)
}

func (s *WhsService) GetBarcodeSuggestion(ctx context.Context, text string, limit int) ([]model.Suggestion, error) {
	return s.barcodesCatalog.Suggest(ctx, text, limit)
}

func (s *WhsService) GetBarcodeTypes(ctx context.Context) ([]model.BarcodeType, error) {
	return s.barcodesCatalog.GetBarcodeTypes(ctx)
}

// UpdateBarcodesByOwner - обновляет штрих-коды у владельца по списку
// если не находим - создаем, находим обновляем, если нужно
func (s *WhsService) UpdateBarcodesByOwner(ctx context.Context, barcodes []model.Barcode, ownerId int64, ownerRef string) error {
	bcItems, err := s.barcodesCatalog.FindByOwnerId(ctx, ownerId, ownerRef)
	if err != nil {
		return err
	}
	bcExistsMap := make(map[string]model.Barcode)
	for i := range bcItems {
		bcExistsMap[bcItems[i].Name] = bcItems[i]
	}

	for i := range barcodes {
		bc, ok := bcExistsMap[barcodes[i].Name]
		if !ok {
			_, err = s.barcodesCatalog.Create(ctx, &barcodes[i])
			if err != nil {
				return err
			}
		} else {
			if bc.Type != barcodes[i].Type {
				_, err = s.barcodesCatalog.Update(ctx, &model.Barcode{
					Id:       bc.Id,
					Name:     barcodes[i].Name,
					Type:     barcodes[i].Type,
					OwnerId:  ownerId,
					OwnerRef: ownerRef,
				})
				return err
			}
		}
	}
	return nil
}
