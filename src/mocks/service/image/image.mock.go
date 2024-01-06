// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/pkg/service/image/image.service.go

// Package mock_image is a generated GoMock package.
package mock_image

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/isd-sgcu/johnjud-gateway/src/app/dto"
	v1 "github.com/isd-sgcu/johnjud-go-proto/johnjud/file/image/v1"
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

// Delete mocks base method.
func (m *MockService) Delete(id string) (bool, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), id)
}

// FindByPetId mocks base method.
func (m *MockService) FindByPetId(id string) ([]*v1.Image, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPetId", id)
	ret0, _ := ret[0].([]*v1.Image)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// FindByPetId indicates an expected call of FindByPetId.
func (mr *MockServiceMockRecorder) FindByPetId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPetId", reflect.TypeOf((*MockService)(nil).FindByPetId), id)
}

// Upload mocks base method.
func (m *MockService) Upload(in *dto.ImageDto) (*v1.Image, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", in)
	ret0, _ := ret[0].(*v1.Image)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockServiceMockRecorder) Upload(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockService)(nil).Upload), in)
}
