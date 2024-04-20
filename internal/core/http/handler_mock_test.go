package http_test

import (
	"github.com/amahrouchi/serena/internal/core/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Route() http.Route {
	args := m.Called()
	return args.Get(0).(http.Route)
}

func (m *MockHandler) Handle(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}
