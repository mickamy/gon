package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/mickamy/gon-example/internal/domain/user/usecase"
)

type ListUserHandler echo.HandlerFunc

func ListUser(uc usecase.ListUser) ListUserHandler {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		output, err := uc.Do(ctx, usecase.ListUserInput{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, output)
	}
}

type GetUserHandler echo.HandlerFunc

func GetUser(uc usecase.GetUser) GetUserHandler {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		output, err := uc.Do(ctx, usecase.GetUserInput{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, output)
	}
}

type CreateUserHandler echo.HandlerFunc

func CreateUser(uc usecase.CreateUser) CreateUserHandler {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		output, err := uc.Do(ctx, usecase.CreateUserInput{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, output)
	}
}

type UpdateUserHandler echo.HandlerFunc

func UpdateUser(uc usecase.UpdateUser) UpdateUserHandler {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		output, err := uc.Do(ctx, usecase.UpdateUserInput{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, output)
	}
}

type DeleteUserHandler echo.HandlerFunc

func DeleteUser(uc usecase.DeleteUser) DeleteUserHandler {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		output, err := uc.Do(ctx, usecase.DeleteUserInput{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, output)
	}
}
