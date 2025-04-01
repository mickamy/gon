package usecase

import (
	"context"
)

type UpdateUserInput struct {
}

type UpdateUserOutput struct {
}

//go:generate mockgen -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
type UpdateUser interface {
	Do(ctx context.Context, input UpdateUserInput) (UpdateUserOutput, error)
}

type updateUser struct {
}

func NewUpdateUser(
) UpdateUser {
	return &updateUser{
	}
}

func (uc updateUser) Do(ctx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	return UpdateUserOutput{}, nil
}
