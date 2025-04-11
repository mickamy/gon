package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mickamy/gon-example/internal/domain/user/usecase"
)

func TestListUsers_Do(t *testing.T) {
	t.Parallel()

	// arrange
	ctx := t.Context()

	// act
	sut := usecase.NewListUsers()
	got, err := sut.Do(ctx, usecase.ListUsersInput{})

	// assert
	require.NoError(t, err)
	assert.Equal(t, usecase.ListUsersOutput{}, got)
}
