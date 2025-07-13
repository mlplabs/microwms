package service

import (
	"context"
	"github.com/mlplabs/mwms-core/whs/model"
)

func (s *WhsService) GetStockData(ctx context.Context, filter int64) (*model.StockData, error) {
	return s.storage.ReportStocks(ctx)
}
