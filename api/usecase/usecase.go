package usecase

import (
	"cash/sample-cash/api/entity"
	"cash/sample-cash/api/infra"
	"context"
)

type Usecase interface {
	GetUser(ctx context.Context, userID string) (entity.User, error)
}

type usecase struct {
	sh infra.SqlHandler
}

func NewUsecase(sh infra.SqlHandler) Usecase {
	return &usecase{sh}
}

func (u *usecase) GetUser(ctx context.Context, userID string) (entity.User, error) {
	user, err := u.sh.GetUser(ctx, userID)
	if err != nil {
		return entity.User{}, err
	}
	return *user, nil
}
