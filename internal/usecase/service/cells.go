package service

import (
	"context"
	"github.com/mlplabs/mwms-core/whs/model"
)

func (s *WhsService) GetCellSuggestion(ctx context.Context, text string, limit int) ([]model.Suggestion, error) {
	return s.cellsCatalog.Suggest(ctx, text, limit)
}
