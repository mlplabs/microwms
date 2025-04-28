package service

import (
	"context"
	"github.com/mlplabs/mwms-core/whs/model"
)

func (s *WhsService) GetUserById(ctx context.Context, userId int64) (*model.User, error) {
	return s.usersCatalog.GetById(ctx, userId)
}

func (s *WhsService) GetUsers(ctx context.Context, offset int, limit int) ([]model.User, int64, error) {
	return s.usersCatalog.GetItems(ctx, offset, limit)
}

func (s *WhsService) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	return s.usersCatalog.Create(ctx, user)
}

func (s *WhsService) UpdateUser(ctx context.Context, user *model.User) (int64, error) {
	return s.usersCatalog.Update(ctx, user)
}

func (s *WhsService) DeleteUser(ctx context.Context, userId int64) error {
	_, err := s.usersCatalog.GetById(ctx, userId)
	if err != nil {
		return err
	}
	return s.usersCatalog.Delete(ctx, userId)
}

func (s *WhsService) GetUserSuggestion(ctx context.Context, text string, limit int) ([]model.Suggestion, error) {
	return s.usersCatalog.Suggest(ctx, text, limit)
}
