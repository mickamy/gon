package repository_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/mickamy/gon-example/internal/domain/user/fixture"
	"github.com/mickamy/gon-example/internal/domain/user/model"
	"github.com/mickamy/gon-example/internal/domain/user/repository"
	"github.com/mickamy/gon-example/internal/infra/storage/database"
)

func TestUser_List(t *testing.T) {
	t.Parallel()

	// arrange
	ctx := t.Context()
	db := initDB(t)
	ms := []model.User{
		fixture.User(nil),
		fixture.User(nil),
		fixture.User(nil),
	}
	require.NoError(t, db.WithContext(ctx).Create(&ms).Error)

	// act
	sut := repository.NewUser(db)
	got, err := sut.List(ctx)

	// assert
	require.NoError(t, err)
	require.Len(t, got, len(ms))
}

func TestUser_Get(t *testing.T) {
	t.Parallel()

	// arrange
	ctx := t.Context()
	db := initDB(t)
	m := fixture.User(nil)
	require.NoError(t, db.WithContext(ctx).Create(&m).Error)

	// act
	sut := repository.NewUser(db)
	got, err := sut.Get(ctx, m.ID)

	// assert
	require.NoError(t, err)
	require.Equal(t, m.ID, got.ID)
}

func TestUser_Find(t *testing.T) {
	t.Parallel()

	// arrange
	ctx := t.Context()
	db := initDB(t)
	m := fixture.User(nil)
	require.NoError(t, db.WithContext(ctx).Create(&m).Error)

	// act
	sut := repository.NewUser(db)
	got, err := sut.Find(ctx, m.ID)

	// assert
	require.NoError(t, err)
	require.Equal(t, m.ID, got.ID)
}

func TestUser_Create(t *testing.T) {
	t.Parallel()

	// arrange
	ctx := t.Context()
	db := initDB(t)
	m := fixture.User(nil)

	// act
	sut := repository.NewUser(db)
	err := sut.Create(ctx, &m)

	// assert
	require.NoError(t, err)
	require.NotEmpty(t, m.ID)
}

func TestUser_Update(t *testing.T) {
	t.Parallel()

	// arrange
	ctx := t.Context()
	db := initDB(t)
	m := fixture.User(nil)
	require.NoError(t, db.WithContext(ctx).Create(&m).Error)

	// act
	sut := repository.NewUser(db)
	m.Name = "Updated Name"
	err := sut.Update(ctx, &m)

	// assert
	require.NoError(t, err)
	var got model.User
	require.NoError(t, db.WithContext(ctx).First(&got, m.ID).Error)
	require.Equal(t, m.ID, got.ID)
	require.Equal(t, m.Name, got.Name)
}

func TestUser_Delete(t *testing.T) {
	t.Parallel()

	// arrange
	ctx := t.Context()
	db := initDB(t)
	m := fixture.User(nil)
	require.NoError(t, db.WithContext(ctx).Create(&m).Error)

	// act
	sut := repository.NewUser(db)
	err := sut.Delete(ctx, m.ID)

	// assert
	require.NoError(t, err)
	var got model.User
	err = db.WithContext(ctx).First(&got, m.ID).Error
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func initDB(t *testing.T) *database.DB {
	// Initialize the database connection here
	// This is a placeholder, replace with actual DB initialization code
	return &database.DB{}
}
