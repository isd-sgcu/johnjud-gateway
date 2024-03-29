// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/pkg/service/image/image.service.go
//
// Generated by this command:
//
//	mockgen -source ./src/pkg/service/image/image.service.go -destination ./src/mocks/service/image/image.mock.go
//

// Package mock_image is a generated GoMock package.
package mock_image

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

// AssignPet mocks base method.
func (m *MockService) AssignPet(arg0 *dto.AssignPetRequest) (*dto.AssignPetResponse, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignPet", arg0)
	ret0, _ := ret[0].(*dto.AssignPetResponse)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// AssignPet indicates an expected call of AssignPet.
func (mr *MockServiceMockRecorder) AssignPet(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignPet", reflect.TypeOf((*MockService)(nil).AssignPet), arg0)
}

// Delete mocks base method.
func (m *MockService) Delete(arg0 string) (*dto.DeleteImageResponse, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*dto.DeleteImageResponse)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), arg0)
}

// FindByPetId mocks base method.
func (m *MockService) FindByPetId(arg0 string) ([]*dto.ImageResponse, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPetId", arg0)
	ret0, _ := ret[0].([]*dto.ImageResponse)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// FindByPetId indicates an expected call of FindByPetId.
func (mr *MockServiceMockRecorder) FindByPetId(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPetId", reflect.TypeOf((*MockService)(nil).FindByPetId), arg0)
}

// Upload mocks base method.
func (m *MockService) Upload(arg0 *dto.UploadImageRequest) (*dto.ImageResponse, *dto.ResponseErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0)
	ret0, _ := ret[0].(*dto.ImageResponse)
	ret1, _ := ret[1].(*dto.ResponseErr)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockServiceMockRecorder) Upload(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockService)(nil).Upload), arg0)
}
