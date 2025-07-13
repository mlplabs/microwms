package service

import (
	"context"
	"github.com/mlplabs/mwms-core/whs/model"
)

func (s *WhsService) GetBarcodes(ctx context.Context, offset int, limit int) ([]model.Barcode, int64, error) {
	return s.storage.GetBarcodesItems(ctx, offset, limit)
}

func (s *WhsService) GetBarcodesByOwner(ctx context.Context, offset int, limit int, ownerId int64, ownerRef string) ([]model.Barcode, int64, error) {
	return s.storage.GetBarcodesItemsByOwner(ctx, offset, limit, ownerId, ownerRef)
}

func (s *WhsService) GetBarcodeById(ctx context.Context, itemId int64) (*model.Barcode, error) {
	return s.storage.GetBarcodeById(ctx, itemId)
}

func (s *WhsService) CreateBarcode(ctx context.Context, bc *model.Barcode) (int64, error) {
	return s.storage.CreateBarcode(ctx, bc)
}

func (s *WhsService) UpdateBarcode(ctx context.Context, bc *model.Barcode) (int64, error) {
	return s.storage.UpdateBarcode(ctx, bc)
}

func (s *WhsService) DeleteBarcode(ctx context.Context, itemId int64) error {
	_, err := s.storage.GetBarcodeById(ctx, itemId)
	if err != nil {
		return err
	}
	return s.storage.DeleteBarcode(ctx, itemId)
}

func (s *WhsService) GetBarcodeSuggestion(ctx context.Context, text string, limit int) ([]model.Suggestion, error) {
	return s.storage.BarcodesSuggest(ctx, text, limit)
}

func (s *WhsService) GetBarcodeTypes(ctx context.Context) ([]model.BarcodeType, error) {
	return s.storage.GetBarcodeTypes(ctx)
}

// UpdateBarcodesByOwner - обновляет штрих-коды у владельца по списку
// если не находим - создаем, находим обновляем, если нужно
func (s *WhsService) UpdateBarcodesByOwner(ctx context.Context, barcodes []model.Barcode, ownerId int64, ownerRef string) error {
	bcItems, err := s.storage.FindBarcodesByOwnerId(ctx, ownerId, ownerRef)
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
			_, err = s.storage.CreateBarcode(ctx, &barcodes[i])
			if err != nil {
				return err
			}
		} else {
			if bc.Type != barcodes[i].Type {
				_, err = s.storage.UpdateBarcode(ctx, &model.Barcode{
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
