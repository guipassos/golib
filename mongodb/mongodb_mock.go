// Code generated by MockGen. DO NOT EDIT.
// Source: mongodb.go

// Package mongodb is a generated GoMock package.
package mongodb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// Collection mocks base method.
func (m *MockDatabase) Collection(name string) Collection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collection", name)
	ret0, _ := ret[0].(Collection)
	return ret0
}

// Collection indicates an expected call of Collection.
func (mr *MockDatabaseMockRecorder) Collection(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collection", reflect.TypeOf((*MockDatabase)(nil).Collection), name)
}

// Name mocks base method.
func (m *MockDatabase) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockDatabaseMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockDatabase)(nil).Name))
}

// URI mocks base method.
func (m *MockDatabase) URI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URI")
	ret0, _ := ret[0].(string)
	return ret0
}

// URI indicates an expected call of URI.
func (mr *MockDatabaseMockRecorder) URI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URI", reflect.TypeOf((*MockDatabase)(nil).URI))
}

// MockCollection is a mock of Collection interface.
type MockCollection struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionMockRecorder
}

// MockCollectionMockRecorder is the mock recorder for MockCollection.
type MockCollectionMockRecorder struct {
	mock *MockCollection
}

// NewMockCollection creates a new mock instance.
func NewMockCollection(ctrl *gomock.Controller) *MockCollection {
	mock := &MockCollection{ctrl: ctrl}
	mock.recorder = &MockCollectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCollection) EXPECT() *MockCollectionMockRecorder {
	return m.recorder
}

// Aggregate mocks base method.
func (m *MockCollection) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, pipeline}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Aggregate", varargs...)
	ret0, _ := ret[0].(*mongo.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Aggregate indicates an expected call of Aggregate.
func (mr *MockCollectionMockRecorder) Aggregate(ctx, pipeline interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, pipeline}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aggregate", reflect.TypeOf((*MockCollection)(nil).Aggregate), varargs...)
}

// BulkWrite mocks base method.
func (m *MockCollection) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, models}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BulkWrite", varargs...)
	ret0, _ := ret[0].(*mongo.BulkWriteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkWrite indicates an expected call of BulkWrite.
func (mr *MockCollectionMockRecorder) BulkWrite(ctx, models interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, models}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkWrite", reflect.TypeOf((*MockCollection)(nil).BulkWrite), varargs...)
}

// CountDocuments mocks base method.
func (m *MockCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountDocuments", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountDocuments indicates an expected call of CountDocuments.
func (mr *MockCollectionMockRecorder) CountDocuments(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountDocuments", reflect.TypeOf((*MockCollection)(nil).CountDocuments), varargs...)
}

// DeleteMany mocks base method.
func (m *MockCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMany", varargs...)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMany indicates an expected call of DeleteMany.
func (mr *MockCollectionMockRecorder) DeleteMany(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMany", reflect.TypeOf((*MockCollection)(nil).DeleteMany), varargs...)
}

// DeleteOne mocks base method.
func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteOne", varargs...)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockCollectionMockRecorder) DeleteOne(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockCollection)(nil).DeleteOne), varargs...)
}

// Distinct mocks base method.
func (m *MockCollection) Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, fieldName, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Distinct", varargs...)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Distinct indicates an expected call of Distinct.
func (mr *MockCollectionMockRecorder) Distinct(ctx, fieldName, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, fieldName, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Distinct", reflect.TypeOf((*MockCollection)(nil).Distinct), varargs...)
}

// Drop mocks base method.
func (m *MockCollection) Drop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Drop indicates an expected call of Drop.
func (mr *MockCollectionMockRecorder) Drop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drop", reflect.TypeOf((*MockCollection)(nil).Drop), ctx)
}

// EstimatedDocumentCount mocks base method.
func (m *MockCollection) EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EstimatedDocumentCount", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimatedDocumentCount indicates an expected call of EstimatedDocumentCount.
func (mr *MockCollectionMockRecorder) EstimatedDocumentCount(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimatedDocumentCount", reflect.TypeOf((*MockCollection)(nil).EstimatedDocumentCount), varargs...)
}

// Find mocks base method.
func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*mongo.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockCollectionMockRecorder) Find(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCollection)(nil).Find), varargs...)
}

// FindOne mocks base method.
func (m *MockCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(*mongo.SingleResult)
	return ret0
}

// FindOne indicates an expected call of FindOne.
func (mr *MockCollectionMockRecorder) FindOne(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockCollection)(nil).FindOne), varargs...)
}

// FindOneAndDelete mocks base method.
func (m *MockCollection) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOneAndDelete", varargs...)
	ret0, _ := ret[0].(*mongo.SingleResult)
	return ret0
}

// FindOneAndDelete indicates an expected call of FindOneAndDelete.
func (mr *MockCollectionMockRecorder) FindOneAndDelete(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneAndDelete", reflect.TypeOf((*MockCollection)(nil).FindOneAndDelete), varargs...)
}

// FindOneAndReplace mocks base method.
func (m *MockCollection) FindOneAndReplace(ctx context.Context, filter, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter, replacement}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOneAndReplace", varargs...)
	ret0, _ := ret[0].(*mongo.SingleResult)
	return ret0
}

// FindOneAndReplace indicates an expected call of FindOneAndReplace.
func (mr *MockCollectionMockRecorder) FindOneAndReplace(ctx, filter, replacement interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter, replacement}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneAndReplace", reflect.TypeOf((*MockCollection)(nil).FindOneAndReplace), varargs...)
}

// FindOneAndUpdate mocks base method.
func (m *MockCollection) FindOneAndUpdate(ctx context.Context, filter, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter, update}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOneAndUpdate", varargs...)
	ret0, _ := ret[0].(*mongo.SingleResult)
	return ret0
}

// FindOneAndUpdate indicates an expected call of FindOneAndUpdate.
func (mr *MockCollectionMockRecorder) FindOneAndUpdate(ctx, filter, update interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter, update}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneAndUpdate", reflect.TypeOf((*MockCollection)(nil).FindOneAndUpdate), varargs...)
}

// Indexes mocks base method.
func (m *MockCollection) Indexes() mongo.IndexView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Indexes")
	ret0, _ := ret[0].(mongo.IndexView)
	return ret0
}

// Indexes indicates an expected call of Indexes.
func (mr *MockCollectionMockRecorder) Indexes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Indexes", reflect.TypeOf((*MockCollection)(nil).Indexes))
}

// InsertMany mocks base method.
func (m *MockCollection) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, documents}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertMany", varargs...)
	ret0, _ := ret[0].(*mongo.InsertManyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMany indicates an expected call of InsertMany.
func (mr *MockCollectionMockRecorder) InsertMany(ctx, documents interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, documents}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMany", reflect.TypeOf((*MockCollection)(nil).InsertMany), varargs...)
}

// InsertOne mocks base method.
func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, document}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertOne", varargs...)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockCollectionMockRecorder) InsertOne(ctx, document interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, document}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockCollection)(nil).InsertOne), varargs...)
}

// Name mocks base method.
func (m *MockCollection) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockCollectionMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockCollection)(nil).Name))
}

// ReplaceOne mocks base method.
func (m *MockCollection) ReplaceOne(ctx context.Context, filter, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter, replacement}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReplaceOne", varargs...)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplaceOne indicates an expected call of ReplaceOne.
func (mr *MockCollectionMockRecorder) ReplaceOne(ctx, filter, replacement interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter, replacement}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceOne", reflect.TypeOf((*MockCollection)(nil).ReplaceOne), varargs...)
}

// UpdateByID mocks base method.
func (m *MockCollection) UpdateByID(ctx context.Context, id, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id, update}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateByID", varargs...)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockCollectionMockRecorder) UpdateByID(ctx, id, update interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id, update}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockCollection)(nil).UpdateByID), varargs...)
}

// UpdateMany mocks base method.
func (m *MockCollection) UpdateMany(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter, update}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMany", varargs...)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMany indicates an expected call of UpdateMany.
func (mr *MockCollectionMockRecorder) UpdateMany(ctx, filter, update interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter, update}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMany", reflect.TypeOf((*MockCollection)(nil).UpdateMany), varargs...)
}

// UpdateOne mocks base method.
func (m *MockCollection) UpdateOne(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter, update}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOne", varargs...)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockCollectionMockRecorder) UpdateOne(ctx, filter, update interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter, update}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockCollection)(nil).UpdateOne), varargs...)
}

// Watch mocks base method.
func (m *MockCollection) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, pipeline}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Watch", varargs...)
	ret0, _ := ret[0].(*mongo.ChangeStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockCollectionMockRecorder) Watch(ctx, pipeline interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, pipeline}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockCollection)(nil).Watch), varargs...)
}
