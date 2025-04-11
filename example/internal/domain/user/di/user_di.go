package di

import (
	"github.com/google/wire"

	"github.com/mickamy/gon-example/internal/domain/user/handler"
	"github.com/mickamy/gon-example/internal/domain/user/repository"
	"github.com/mickamy/gon-example/internal/domain/user/usecase"
)

type Repositories struct {
	repository.User
}

//lint:ignore U1000 used by wire
var RepositorySet = wire.NewSet(
	repository.NewUser,
)

type UseCases struct {
	usecase.CreateUser
	usecase.DeleteUser
	usecase.GetUser
	usecase.ListUser
	usecase.UpdateUser
}

//lint:ignore U1000 used by wire
var UseCaseSet = wire.NewSet(
	usecase.NewCreateUser,
	usecase.NewDeleteUser,
	usecase.NewGetUser,
	usecase.NewListUser,
	usecase.NewUpdateUser,
)

type Handlers struct {
	*handler.ListUserHandler
	*handler.GetUserHandler
	*handler.CreateUserHandler
	*handler.UpdateUserHandler
	*handler.DeleteUserHandler
}

//lint:ignore U1000 used by wire
var HandlerSet = wire.NewSet(
	handler.ListUser,
	handler.GetUser,
	handler.CreateUser,
	handler.UpdateUser,
	handler.DeleteUser,
)
