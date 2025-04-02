package handler

import (
	"net/http"
)

type ListUserHandler http.HandlerFunc

func ListUser() ListUserHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
	}
}

type GetUserHandler http.HandlerFunc

func GetUser() GetUserHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
	}
}

type CreateUserHandler http.HandlerFunc

func CreateUser() CreateUserHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
	}
}

type UpdateUserHandler http.HandlerFunc

func UpdateUser() UpdateUserHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
	}
}

type DeleteUserHandler http.HandlerFunc

func DeleteUser() DeleteUserHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
	}
}
