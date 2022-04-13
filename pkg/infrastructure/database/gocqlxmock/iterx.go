package gocqlxmock

import (
	"reflect"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/mock"
)

type IIterx interface {
	Unsafe() IIterx
	StructOnly() IIterx
	Get(dest interface{}) error
	scanAny(dest interface{}) bool
	Select(dest interface{}) error
	scanAll(dest interface{}) bool
	isScannable(t reflect.Type) bool
	scan(value reflect.Value) bool
	StructScan(dest interface{}) bool
	structScan(value reflect.Value) bool
	fieldsByTraversal(value reflect.Value, traversals [][]int, values []interface{}) error
	Scan(dest ...interface{}) bool
	Close() error
	checkErrAndNotFound() error
}

type IterxMock struct {
	mock.Mock
}

func (mock IterxMock) Unsafe() IIterx {
	args := mock.Called()

	return args.Get(0).(IIterx)
}

func (mock IterxMock) StructOnly() IIterx {
	args := mock.Called()

	return args.Get(0).(IIterx)
}

func (mock IterxMock) Get(dest interface{}) error {
	args := mock.Called(dest)

	return args.Error(0)
}

func (mock IterxMock) scanAny(dest interface{}) bool {
	args := mock.Called(dest)

	return args.Get(0).(bool)
}

func (mock IterxMock) Select(dest interface{}) error {
	args := mock.Called(dest)

	return args.Error(0)
}

func (mock IterxMock) scanAll(dest interface{}) bool {
	args := mock.Called(dest)

	return args.Get(0).(bool)
}

func (mock IterxMock) isScannable(t reflect.Type) bool {
	args := mock.Called(t)

	return args.Get(0).(bool)
}

func (mock IterxMock) scan(value reflect.Value) bool {
	args := mock.Called(value)

	return args.Get(0).(bool)
}

func (mock IterxMock) StructScan(dest interface{}) bool {
	args := mock.Called(dest)

	return args.Get(0).(bool)
}

func (mock IterxMock) structScan(value reflect.Value) bool {
	args := mock.Called(value)

	return args.Get(0).(bool)
}

func (mock IterxMock) fieldsByTraversal(value reflect.Value, traversals [][]int, values []interface{}) error {
	args := mock.Called(value, traversals, values)

	return args.Error(0)
}

func (mock IterxMock) Scan(dest ...interface{}) bool {
	args := mock.Called(dest)

	return args.Get(0).(bool)
}

func (mock IterxMock) Close() error {
	args := mock.Called()

	return args.Error(0)
}

func (mock IterxMock) checkErrAndNotFound() error {
	args := mock.Called()

	return args.Error(0)
}

// "Interface assertion"
var (
	_ IIterx = IterxMock{}
	_ IIterx = iterx{}
)

type iterx struct {
	i *gocql.Iter
}

func (i iterx) Unsafe() IIterx {
	return i
}

func (i iterx) StructOnly() IIterx {
	return i
}

func (i iterx) Get(dest interface{}) error {
	return nil
}

func (i iterx) scanAny(dest interface{}) bool {
	return true
}

func (i iterx) Select(dest interface{}) error {
	return nil
}

func (i iterx) scanAll(dest interface{}) bool {
	return true
}

func (i iterx) isScannable(t reflect.Type) bool {
	return true
}

func (i iterx) scan(value reflect.Value) bool {
	return true
}

func (i iterx) StructScan(dest interface{}) bool {
	return true
}

func (i iterx) structScan(value reflect.Value) bool {
	return true
}

func (i iterx) fieldsByTraversal(value reflect.Value, traversals [][]int, values []interface{}) error {
	return nil
}

func (i iterx) Scan(dest ...interface{}) bool {
	return true
}

func (i iterx) Close() error {
	return nil
}

func (i iterx) checkErrAndNotFound() error {
	return nil
}
