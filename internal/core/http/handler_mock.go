package http

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type mockHandler struct {
	mock.Mock
}

func (m *mockHandler) Route() Route {
	args := m.Called()
	return args.Get(0).(Route)
}

func (m *mockHandler) Handle(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}
