// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/pkg/service/pet/pet.service.go
//
// Generated by this command:
//
//	mockgen -source ./src/pkg/service/pet/pet.service.go -destination ./src/mocks/service/pet/pet.mock.go
//

// Package mock_pet is a generated GoMock package.
package mock_pet

import (
	reflect "reflect"

	dto "github.com/isd-sgcu/johnjud-gateway/src/app/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Adopt mocks base method.
func (m *MockService) Adopt(arg0 string, arg1 *dto.AdoptByRequest) (*dto.AdoptByResponse, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Adopt", arg0, arg1)
	ret0, _ := ret[0].(*dto.AdoptByResponse)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// Adopt indicates an expected call of Adopt.
func (mr *MockServiceMockRecorder) Adopt(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Adopt", reflect.TypeOf((*MockService)(nil).Adopt), arg0, arg1)
}

// ChangeView mocks base method.
func (m *MockService) ChangeView(arg0 string, arg1 *dto.ChangeViewPetRequest) (*dto.ChangeViewPetResponse, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeView", arg0, arg1)
	ret0, _ := ret[0].(*dto.ChangeViewPetResponse)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// ChangeView indicates an expected call of ChangeView.
func (mr *MockServiceMockRecorder) ChangeView(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeView", reflect.TypeOf((*MockService)(nil).ChangeView), arg0, arg1)
}

// Create mocks base method.
func (m *MockService) Create(arg0 *dto.CreatePetRequest) (*dto.PetResponse, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*dto.PetResponse)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockService) Delete(arg0 string) (*dto.DeleteResponse, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*dto.DeleteResponse)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), arg0)
}

// FindAll mocks base method.
func (m *MockService) FindAll() ([]*dto.PetResponse, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*dto.PetResponse)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockServiceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockService)(nil).FindAll))
}

// FindOne mocks base method.
func (m *MockService) FindOne(arg0 string) (*dto.PetResponse, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0)
	ret0, _ := ret[0].(*dto.PetResponse)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockServiceMockRecorder) FindOne(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockService)(nil).FindOne), arg0)
}

// Update mocks base method.
func (m *MockService) Update(arg0 string, arg1 *dto.UpdatePetRequest) (*dto.PetResponse, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*dto.PetResponse)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), arg0, arg1)
}
