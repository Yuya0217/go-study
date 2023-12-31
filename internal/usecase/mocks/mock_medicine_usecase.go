// Code generated by MockGen. DO NOT EDIT.
// Source: medicine_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	dtos "go-layered-architecture-sample/internal/usecase/dtos"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMedicineUseCase is a mock of MedicineUseCase interface.
type MockMedicineUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockMedicineUseCaseMockRecorder
}

// MockMedicineUseCaseMockRecorder is the mock recorder for MockMedicineUseCase.
type MockMedicineUseCaseMockRecorder struct {
	mock *MockMedicineUseCase
}

// NewMockMedicineUseCase creates a new mock instance.
func NewMockMedicineUseCase(ctrl *gomock.Controller) *MockMedicineUseCase {
	mock := &MockMedicineUseCase{ctrl: ctrl}
	mock.recorder = &MockMedicineUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMedicineUseCase) EXPECT() *MockMedicineUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMedicineUseCase) Create(ctx context.Context, input dtos.Medicine) (*dtos.Medicine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(*dtos.Medicine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMedicineUseCaseMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMedicineUseCase)(nil).Create), ctx, input)
}

// Delete mocks base method.
func (m *MockMedicineUseCase) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMedicineUseCaseMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMedicineUseCase)(nil).Delete), ctx, id)
}

// Find mocks base method.
func (m *MockMedicineUseCase) Find(ctx context.Context, params dtos.MedicineFindParams) ([]*dtos.Medicine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, params)
	ret0, _ := ret[0].([]*dtos.Medicine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMedicineUseCaseMockRecorder) Find(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMedicineUseCase)(nil).Find), ctx, params)
}

// GetByID mocks base method.
func (m *MockMedicineUseCase) GetByID(ctx context.Context, id int) (*dtos.Medicine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*dtos.Medicine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockMedicineUseCaseMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMedicineUseCase)(nil).GetByID), ctx, id)
}

// Update mocks base method.
func (m *MockMedicineUseCase) Update(ctx context.Context, input dtos.Medicine) (*dtos.Medicine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, input)
	ret0, _ := ret[0].(*dtos.Medicine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMedicineUseCaseMockRecorder) Update(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMedicineUseCase)(nil).Update), ctx, input)
}
