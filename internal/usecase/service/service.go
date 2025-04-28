package service

import (
	"github.com/mlplabs/mwms-core/whs"
)

type WhsService struct {
	wms                  *whs.Wms
	usersCatalog         *whs.Users
	barcodesCatalog      *whs.Barcodes
	manufacturersCatalog *whs.Manufacturers
	productsCatalog      *whs.Products
	warehousesCatalog    *whs.Warehouses
	cellsCatalog         *whs.Cells
	reports              *whs.Reports
	storage              *whs.Storage
}

func NewWhsService(wms *whs.Wms) *WhsService {
	return &WhsService{
		wms:                  wms,
		usersCatalog:         whs.NewUsers(wms),
		barcodesCatalog:      whs.NewBarcodes(wms),
		manufacturersCatalog: whs.NewManufacturers(wms),
		productsCatalog:      whs.NewProducts(wms),
		warehousesCatalog:    whs.NewWarehouses(wms),
		cellsCatalog:         whs.NewCells(wms),
		reports:              whs.NewReports(wms),
		storage:              whs.NewStorage(wms),
	}
}
