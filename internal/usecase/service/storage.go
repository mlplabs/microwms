package service

import "context"

func (s *WhsService) GetItemFromCell(ctx context.Context, prodId int64, cellId int64, quantity int) (int, error) {
	return s.storage.GetItemFromCell(ctx, prodId, cellId, quantity)
}

func (s *WhsService) PutItemToCell(ctx context.Context, prodId int64, cellId int64, quantity int) (int, error) {
	return s.storage.PutItemToCell(ctx, prodId, cellId, quantity)
}

func (s *WhsService) MoveItemToCell(ctx context.Context, prodId int64, cellSrcId int64, cellDstId int64, quantity int) (int, error) {
	return s.storage.MoveItemToCell(ctx, prodId, cellSrcId, cellDstId, quantity)
}
