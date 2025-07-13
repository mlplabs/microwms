package service

import (
	"context"
	"github.com/mlplabs/mwms-core/whs/model"
)

func (s *WhsService) GetUserById(ctx context.Context, userId int64) (*model.User, error) {
	return s.storage.GetUserById(ctx, userId)
}

func (s *WhsService) GetUsers(ctx context.Context, offset int, limit int) ([]model.User, int64, error) {
	return s.storage.GetUsersItems(ctx, offset, limit)
}

func (s *WhsService) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	return s.storage.CreateUser(ctx, user)
}

func (s *WhsService) UpdateUser(ctx context.Context, user *model.User) (int64, error) {
	return s.storage.UpdateUser(ctx, user)
}

func (s *WhsService) DeleteUser(ctx context.Context, userId int64) error {
	_, err := s.storage.GetUserById(ctx, userId)
	if err != nil {
		return err
	}
	return s.storage.DeleteUser(ctx, userId)
}

func (s *WhsService) GetUserSuggestion(ctx context.Context, text string, limit int) ([]model.Suggestion, error) {
	return s.storage.UsersSuggest(ctx, text, limit)
}
