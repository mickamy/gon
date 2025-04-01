package database

import (
	"errors"

	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("record not found")

// DB is a wrapper of *gorm.DB
type DB struct{ *gorm.DB }
