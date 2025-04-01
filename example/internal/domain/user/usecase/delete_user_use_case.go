package usecase

import (
	"context"
)

type DeleteUserInput struct {
}

type DeleteUserOutput struct {
}

//go:generate mockgen -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
type DeleteUser interface {
	Do(ctx context.Context, input DeleteUserInput) (DeleteUserOutput, error)
}

type deleteUser struct {
}

func NewDeleteUser(
) DeleteUser {
	return &deleteUser{
	}
}

func (uc deleteUser) Do(ctx context.Context, input DeleteUserInput) (DeleteUserOutput, error) {
	return DeleteUserOutput{}, nil
}
