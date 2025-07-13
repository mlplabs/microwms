package service

import (
	"github.com/mlplabs/mwms-core/whs"
)

type WhsService struct {
	wms     *whs.Wms
	storage *whs.Storage
}

func NewWhsService(wms *whs.Wms) *WhsService {
	return &WhsService{
		wms:     wms,
		storage: whs.NewStorage(wms),
	}
}
