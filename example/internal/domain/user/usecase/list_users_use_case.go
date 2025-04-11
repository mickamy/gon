package usecase

import (
	"context"
)

type ListUsersInput struct {
}

type ListUsersOutput struct {
}

//go:generate mockgen -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
type ListUsers interface {
	Do(ctx context.Context, input ListUsersInput) (ListUsersOutput, error)
}

type listUsers struct {
}

func NewListUsers() ListUsers {
	return &listUsers{}
}

func (uc listUsers) Do(ctx context.Context, input ListUsersInput) (ListUsersOutput, error) {
	return ListUsersOutput{}, nil
}
