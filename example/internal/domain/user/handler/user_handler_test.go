package handler_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mickamy/gon-example/internal/domain/user/handler"
	"github.com/mickamy/gon-example/internal/domain/user/usecase"
	"github.com/mickamy/gon-example/test/httptestutil"
)

func TestListUser(t *testing.T) {
	t.Parallel()

	// arrange
	c, recorder := httptestutil.NewEchoRequest(t, httptestutil.RequestBuilder{
		Method:      http.MethodGet,
		Path:        "/user",
		Query:       nil,
		Body:        nil,
		ContentType: echo.MIMEApplicationJSON,
	})
	uc := usecase.NewListUser()

	// act
	sut := handler.ListUser(uc)
	require.NoError(t, sut(c))

	// assert
	require.Equal(t, http.StatusOK, recorder.Code)
	var res interface{}
	assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&res))
	assert.NotNil(t, res)
}

func TestGetUser(t *testing.T) {
	t.Parallel()

	// arrange
	c, recorder := httptestutil.NewEchoRequest(t, httptestutil.RequestBuilder{
		Method:      http.MethodGet,
		Path:        "/user",
		Query:       nil,
		Body:        nil,
		ContentType: echo.MIMEApplicationJSON,
	})
	uc := usecase.NewGetUser()

	// act
	sut := handler.GetUser(uc)
	require.NoError(t, sut(c))

	// assert
	require.Equal(t, http.StatusOK, recorder.Code)
	var res interface{}
	assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&res))
	assert.NotNil(t, res)
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	// arrange
	c, recorder := httptestutil.NewEchoRequest(t, httptestutil.RequestBuilder{
		Method:      http.MethodGet,
		Path:        "/user",
		Query:       nil,
		Body:        nil,
		ContentType: echo.MIMEApplicationJSON,
	})
	uc := usecase.NewCreateUser()

	// act
	sut := handler.CreateUser(uc)
	require.NoError(t, sut(c))

	// assert
	require.Equal(t, http.StatusOK, recorder.Code)
	var res interface{}
	assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&res))
	assert.NotNil(t, res)
}

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	// arrange
	c, recorder := httptestutil.NewEchoRequest(t, httptestutil.RequestBuilder{
		Method:      http.MethodGet,
		Path:        "/user",
		Query:       nil,
		Body:        nil,
		ContentType: echo.MIMEApplicationJSON,
	})
	uc := usecase.NewUpdateUser()

	// act
	sut := handler.UpdateUser(uc)
	require.NoError(t, sut(c))

	// assert
	require.Equal(t, http.StatusOK, recorder.Code)
	var res interface{}
	assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&res))
	assert.NotNil(t, res)
}

func TestDeleteUser(t *testing.T) {
	t.Parallel()

	// arrange
	c, recorder := httptestutil.NewEchoRequest(t, httptestutil.RequestBuilder{
		Method:      http.MethodGet,
		Path:        "/user",
		Query:       nil,
		Body:        nil,
		ContentType: echo.MIMEApplicationJSON,
	})
	uc := usecase.NewDeleteUser()

	// act
	sut := handler.DeleteUser(uc)
	require.NoError(t, sut(c))

	// assert
	require.Equal(t, http.StatusOK, recorder.Code)
	var res interface{}
	assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&res))
	assert.NotNil(t, res)
}
