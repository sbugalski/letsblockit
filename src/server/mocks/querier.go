// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/db/querier.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	pgtype "github.com/jackc/pgtype"
	db "github.com/xvello/letsblockit/src/db"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CountInstanceForUserAndFilter mocks base method.
func (m *MockQuerier) CountInstanceForUserAndFilter(ctx context.Context, arg db.CountInstanceForUserAndFilterParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountInstanceForUserAndFilter", ctx, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountInstanceForUserAndFilter indicates an expected call of CountInstanceForUserAndFilter.
func (mr *MockQuerierMockRecorder) CountInstanceForUserAndFilter(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountInstanceForUserAndFilter", reflect.TypeOf((*MockQuerier)(nil).CountInstanceForUserAndFilter), ctx, arg)
}

// CountListsForUser mocks base method.
func (m *MockQuerier) CountListsForUser(ctx context.Context, userID uuid.UUID) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountListsForUser", ctx, userID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountListsForUser indicates an expected call of CountListsForUser.
func (mr *MockQuerierMockRecorder) CountListsForUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountListsForUser", reflect.TypeOf((*MockQuerier)(nil).CountListsForUser), ctx, userID)
}

// CreateInstanceForUserAndFilter mocks base method.
func (m *MockQuerier) CreateInstanceForUserAndFilter(ctx context.Context, arg db.CreateInstanceForUserAndFilterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstanceForUserAndFilter", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInstanceForUserAndFilter indicates an expected call of CreateInstanceForUserAndFilter.
func (mr *MockQuerierMockRecorder) CreateInstanceForUserAndFilter(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstanceForUserAndFilter", reflect.TypeOf((*MockQuerier)(nil).CreateInstanceForUserAndFilter), ctx, arg)
}

// CreateListForUser mocks base method.
func (m *MockQuerier) CreateListForUser(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateListForUser", ctx, userID)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateListForUser indicates an expected call of CreateListForUser.
func (mr *MockQuerierMockRecorder) CreateListForUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateListForUser", reflect.TypeOf((*MockQuerier)(nil).CreateListForUser), ctx, userID)
}

// CreateUser mocks base method.
func (m *MockQuerier) CreateUser(ctx context.Context, arg db.CreateUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockQuerierMockRecorder) CreateUser(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockQuerier)(nil).CreateUser), ctx, arg)
}

// DeleteInstanceForUserAndFilter mocks base method.
func (m *MockQuerier) DeleteInstanceForUserAndFilter(ctx context.Context, arg db.DeleteInstanceForUserAndFilterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInstanceForUserAndFilter", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInstanceForUserAndFilter indicates an expected call of DeleteInstanceForUserAndFilter.
func (mr *MockQuerierMockRecorder) DeleteInstanceForUserAndFilter(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstanceForUserAndFilter", reflect.TypeOf((*MockQuerier)(nil).DeleteInstanceForUserAndFilter), ctx, arg)
}

// GetActiveFiltersForUser mocks base method.
func (m *MockQuerier) GetActiveFiltersForUser(ctx context.Context, userID uuid.UUID) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveFiltersForUser", ctx, userID)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveFiltersForUser indicates an expected call of GetActiveFiltersForUser.
func (mr *MockQuerierMockRecorder) GetActiveFiltersForUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveFiltersForUser", reflect.TypeOf((*MockQuerier)(nil).GetActiveFiltersForUser), ctx, userID)
}

// GetCookieKeys mocks base method.
func (m *MockQuerier) GetCookieKeys(ctx context.Context) (db.GetCookieKeysRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCookieKeys", ctx)
	ret0, _ := ret[0].(db.GetCookieKeysRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCookieKeys indicates an expected call of GetCookieKeys.
func (mr *MockQuerierMockRecorder) GetCookieKeys(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCookieKeys", reflect.TypeOf((*MockQuerier)(nil).GetCookieKeys), ctx)
}

// GetInstanceForUserAndFilter mocks base method.
func (m *MockQuerier) GetInstanceForUserAndFilter(ctx context.Context, arg db.GetInstanceForUserAndFilterParams) (pgtype.JSONB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceForUserAndFilter", ctx, arg)
	ret0, _ := ret[0].(pgtype.JSONB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceForUserAndFilter indicates an expected call of GetInstanceForUserAndFilter.
func (mr *MockQuerierMockRecorder) GetInstanceForUserAndFilter(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceForUserAndFilter", reflect.TypeOf((*MockQuerier)(nil).GetInstanceForUserAndFilter), ctx, arg)
}

// GetInstancesForList mocks base method.
func (m *MockQuerier) GetInstancesForList(ctx context.Context, filterListID int32) ([]db.GetInstancesForListRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstancesForList", ctx, filterListID)
	ret0, _ := ret[0].([]db.GetInstancesForListRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstancesForList indicates an expected call of GetInstancesForList.
func (mr *MockQuerierMockRecorder) GetInstancesForList(ctx, filterListID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstancesForList", reflect.TypeOf((*MockQuerier)(nil).GetInstancesForList), ctx, filterListID)
}

// GetListForToken mocks base method.
func (m *MockQuerier) GetListForToken(ctx context.Context, token uuid.UUID) (db.GetListForTokenRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListForToken", ctx, token)
	ret0, _ := ret[0].(db.GetListForTokenRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListForToken indicates an expected call of GetListForToken.
func (mr *MockQuerierMockRecorder) GetListForToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListForToken", reflect.TypeOf((*MockQuerier)(nil).GetListForToken), ctx, token)
}

// GetListForUser mocks base method.
func (m *MockQuerier) GetListForUser(ctx context.Context, userID uuid.UUID) (db.GetListForUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListForUser", ctx, userID)
	ret0, _ := ret[0].(db.GetListForUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListForUser indicates an expected call of GetListForUser.
func (mr *MockQuerierMockRecorder) GetListForUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListForUser", reflect.TypeOf((*MockQuerier)(nil).GetListForUser), ctx, userID)
}

// GetStats mocks base method.
func (m *MockQuerier) GetStats(ctx context.Context) (db.GetStatsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats", ctx)
	ret0, _ := ret[0].(db.GetStatsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats.
func (mr *MockQuerierMockRecorder) GetStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockQuerier)(nil).GetStats), ctx)
}

// GetUserByEmail mocks base method.
func (m *MockQuerier) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockQuerierMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockQuerier)(nil).GetUserByEmail), ctx, email)
}

// HasUserDownloadedList mocks base method.
func (m *MockQuerier) HasUserDownloadedList(ctx context.Context, userID uuid.UUID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUserDownloadedList", ctx, userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasUserDownloadedList indicates an expected call of HasUserDownloadedList.
func (mr *MockQuerierMockRecorder) HasUserDownloadedList(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUserDownloadedList", reflect.TypeOf((*MockQuerier)(nil).HasUserDownloadedList), ctx, userID)
}

// MarkListDownloaded mocks base method.
func (m *MockQuerier) MarkListDownloaded(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkListDownloaded", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkListDownloaded indicates an expected call of MarkListDownloaded.
func (mr *MockQuerierMockRecorder) MarkListDownloaded(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkListDownloaded", reflect.TypeOf((*MockQuerier)(nil).MarkListDownloaded), ctx, id)
}

// RotateListToken mocks base method.
func (m *MockQuerier) RotateListToken(ctx context.Context, arg db.RotateListTokenParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RotateListToken", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// RotateListToken indicates an expected call of RotateListToken.
func (mr *MockQuerierMockRecorder) RotateListToken(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateListToken", reflect.TypeOf((*MockQuerier)(nil).RotateListToken), ctx, arg)
}

// SetCookieKeys mocks base method.
func (m *MockQuerier) SetCookieKeys(ctx context.Context, arg db.SetCookieKeysParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCookieKeys", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCookieKeys indicates an expected call of SetCookieKeys.
func (mr *MockQuerierMockRecorder) SetCookieKeys(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCookieKeys", reflect.TypeOf((*MockQuerier)(nil).SetCookieKeys), ctx, arg)
}

// UpdateInstanceForUserAndFilter mocks base method.
func (m *MockQuerier) UpdateInstanceForUserAndFilter(ctx context.Context, arg db.UpdateInstanceForUserAndFilterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInstanceForUserAndFilter", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInstanceForUserAndFilter indicates an expected call of UpdateInstanceForUserAndFilter.
func (mr *MockQuerierMockRecorder) UpdateInstanceForUserAndFilter(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstanceForUserAndFilter", reflect.TypeOf((*MockQuerier)(nil).UpdateInstanceForUserAndFilter), ctx, arg)
}

// UpdateUser mocks base method.
func (m *MockQuerier) UpdateUser(ctx context.Context, arg db.UpdateUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockQuerierMockRecorder) UpdateUser(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockQuerier)(nil).UpdateUser), ctx, arg)
}
