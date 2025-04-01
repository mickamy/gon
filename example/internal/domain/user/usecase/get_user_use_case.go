package usecase

import (
	"context"
)

type GetUserInput struct {
}

type GetUserOutput struct {
}

//go:generate mockgen -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
type GetUser interface {
	Do(ctx context.Context, input GetUserInput) (GetUserOutput, error)
}

type getUser struct {
}

func NewGetUser() GetUser {
	return &getUser{}
}

func (uc getUser) Do(ctx context.Context, input GetUserInput) (GetUserOutput, error) {
	return GetUserOutput{}, nil
}
