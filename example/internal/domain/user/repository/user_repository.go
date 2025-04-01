package repository

import (
    "context"
    "errors"

	"gorm.io/gorm"

    "github.com/mickamy/gon-example/internal/domain/user/model"
	"github.com/mickamy/gon-example/internal/infra/storage/database"
)

//go:generate mockgen -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
type User interface {
    List(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]model.User, error)
	Get(ctx context.Context, id string, scopes ...func(*gorm.DB) *gorm.DB) (model.User, error)
	Find(ctx context.Context, id string, scopes ...func(*gorm.DB) *gorm.DB) (*model.User, error)
	Create(ctx context.Context, m *model.User) error
	Update(ctx context.Context, m *model.User) error
	Delete(ctx context.Context, id string) error
	WithTx(tx *database.DB) User
}

type user struct {
    db *database.DB
}

func NewUser(db *database.DB) User {
    return &user{db: db}
}

func (repo *user) List(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]model.User, error) {
    var m []model.User
    err := repo.db.WithContext(ctx).Scopes(scopes...).Find(&m).Error
    return m, err
}

func (repo *user) Get(ctx context.Context, id string, scopes ...func(*gorm.DB) *gorm.DB) (model.User, error) {
    var m model.User
    err := repo.db.WithContext(ctx).Scopes(scopes...).First(&m, id).Error
    return m, err
}

func (repo *user) Find(ctx context.Context, id string, scopes ...func(*gorm.DB) *gorm.DB) (*model.User, error) {
    var m model.User
    if err := repo.db.WithContext(ctx).Scopes(scopes...).First(&m, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &m, nil
}

func (repo *user) Create(ctx context.Context, m *model.User) error {
    return repo.db.WithContext(ctx).Create(&m).Error
}

func (repo *user) Update(ctx context.Context, m *model.User) error {
    return repo.db.WithContext(ctx).Save(m).Error
}

func (repo *user) Delete(ctx context.Context, id string) error {
    return repo.db.WithContext(ctx).Delete(&model.User{}, id).Error
}

func (repo *user) WithTx(tx *database.DB) User {
    return &user{db: tx}
}
