package services

import "github.com/stretchr/testify/mock"

// PayloadWriterMock is a mock for the PayloadWriter service.
type PayloadWriterMock struct {
	mock.Mock
}

func (pwm *PayloadWriterMock) Write(author string, data map[string]interface{}) error {
	args := pwm.Called(author, data)
	return args.Error(0)
}
