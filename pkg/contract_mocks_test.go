// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package tg_bot is a generated GoMock package.
package tg_bot

import (
	reflect "reflect"

	telegram "github.com/Isotere/tg-bot/pkg/internal/clients/telegram"
	gomock "github.com/golang/mock/gomock"
)

// MockTgClient is a mock of TgClient interface.
type MockTgClient struct {
	ctrl     *gomock.Controller
	recorder *MockTgClientMockRecorder
}

// MockTgClientMockRecorder is the mock recorder for MockTgClient.
type MockTgClientMockRecorder struct {
	mock *MockTgClient
}

// NewMockTgClient creates a new mock instance.
func NewMockTgClient(ctrl *gomock.Controller) *MockTgClient {
	mock := &MockTgClient{ctrl: ctrl}
	mock.recorder = &MockTgClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTgClient) EXPECT() *MockTgClientMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockTgClient) SendMessage(chatID int, text string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", chatID, text)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockTgClientMockRecorder) SendMessage(chatID, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockTgClient)(nil).SendMessage), chatID, text)
}

// Updates mocks base method.
func (m *MockTgClient) Updates(offset, limit int) ([]telegram.Update, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", offset, limit)
	ret0, _ := ret[0].([]telegram.Update)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Updates indicates an expected call of Updates.
func (mr *MockTgClientMockRecorder) Updates(offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*MockTgClient)(nil).Updates), offset, limit)
}
