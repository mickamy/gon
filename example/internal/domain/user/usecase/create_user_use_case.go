package usecase

import (
	"context"
)

type CreateUserInput struct {
}

type CreateUserOutput struct {
}

//go:generate mockgen -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
type CreateUser interface {
	Do(ctx context.Context, input CreateUserInput) (CreateUserOutput, error)
}

type createUser struct {
}

func NewCreateUser(
) CreateUser {
	return &createUser{
	}
}

func (uc createUser) Do(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	return CreateUserOutput{}, nil
}
