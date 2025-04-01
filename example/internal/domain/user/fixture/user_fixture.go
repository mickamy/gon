package fixture

import (
	"github.com/mickamy/gon-example/internal/domain/user/model"
)

func User(setter func(m *model.User)) model.User {
	m := model.User{
		// TODO: Set sample values for each field
	}

	if setter != nil {
		setter(&m)
	}

	return m
}
