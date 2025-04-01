package usecase

import (
	"context"
)

type ListUserInput struct {
}

type ListUserOutput struct {
}

//go:generate mockgen -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
type ListUser interface {
	Do(ctx context.Context, input ListUserInput) (ListUserOutput, error)
}

type listUser struct {
}

func NewListUser(
) ListUser {
	return &listUser{
	}
}

func (uc listUser) Do(ctx context.Context, input ListUserInput) (ListUserOutput, error) {
	return ListUserOutput{}, nil
}
