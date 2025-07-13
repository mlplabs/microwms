package service

import (
	"context"
	"github.com/mlplabs/mwms-core/whs/model"
)

const ownerRefProducts = "products"

func (s *WhsService) GetProductById(ctx context.Context, productId int64) (*model.Product, error) {
	product, err := s.storage.GetProductById(ctx, productId)
	if err != nil {
		return nil, err
	}
	barcodes, err := s.storage.FindBarcodesByOwnerId(ctx, productId, ownerRefProducts)
	if err != nil {
		return nil, err
	}
	product.Barcodes = barcodes
	return product, nil
}

func (s *WhsService) GetProducts(ctx context.Context, offset int, limit int, search string) ([]model.Product, int64, error) {
	return s.storage.GetProductsItems(ctx, offset, limit, search)
}
func (s *WhsService) CreateProduct(ctx context.Context, product *model.Product) (int64, error) {
	if product.Manufacturer.Name != "" {
		mnfItem, err := s.FindOrCreateManufacturer(ctx, product.Manufacturer.Name)
		if err != nil {
			return 0, err
		}
		product.Manufacturer = *mnfItem
	} else {
		product.Manufacturer = model.Manufacturer{}
	}

	newItemId, err := s.storage.CreateProduct(ctx, product)
	if err != nil {
		return 0, err
	}

	if len(product.Barcodes) > 0 {
		err = s.UpdateBarcodesByOwner(ctx, product.Barcodes, newItemId, ownerRefProducts) // TODO: !!!
		if err != nil {
			return 0, err
		}
	}
	return newItemId, err
}

func (s *WhsService) UpdateProduct(ctx context.Context, product *model.Product) (int64, error) {
	if product.Manufacturer.Name != "" {
		mnfItem, err := s.FindOrCreateManufacturer(ctx, product.Manufacturer.Name)
		if err != nil {
			return 0, err
		}
		product.Manufacturer = *mnfItem
	} else {
		product.Manufacturer = model.Manufacturer{}
	}

	if len(product.Barcodes) > 0 {
		bcItems, err := s.storage.FindBarcodesByOwnerId(ctx, product.Id, ownerRefProducts)
		if err != nil {
			return 0, err
		}
		bcItemsMap := make(map[string]model.Barcode)
		for i := range bcItems {
			bcItemsMap[bcItems[i].Name] = bcItems[i]
		}

		for j := range product.Barcodes {
			if _, ok := bcItemsMap[product.Barcodes[j].Name]; !ok {
				_, err = s.storage.CreateBarcode(ctx, &model.Barcode{
					Name:     product.Barcodes[j].Name,
					Type:     product.Barcodes[j].Type,
					OwnerId:  product.Id,
					OwnerRef: ownerRefProducts,
				})
				if err != nil {
					return 0, err
				}
			}

			if err != nil {
				return 0, err
			}
		}
	}

	return s.storage.UpdateProduct(ctx, product)
}

func (s *WhsService) DeleteProduct(ctx context.Context, productId int64) error {
	_, err := s.storage.GetProductById(ctx, productId)
	if err != nil {
		return err
	}
	return s.storage.DeleteProduct(ctx, productId)
}

func (s *WhsService) FindProductsByName(ctx context.Context, name string) ([]model.Product, error) {
	return s.storage.FindProductsByName(ctx, name)
}
func (s *WhsService) FindProductsByBarcode(ctx context.Context, name string) ([]model.Product, error) {
	return s.storage.FindProductsByBarcode(ctx, name)
}
func (s *WhsService) GetProductsSuggestion(ctx context.Context, text string, limit int) ([]model.Suggestion, error) {
	return s.storage.ProductsSuggest(ctx, text, limit)
}
