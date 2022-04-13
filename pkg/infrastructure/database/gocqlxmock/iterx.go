package gocqlxmock

import (
	"github.com/Guilospanck/igocqlx"
	"github.com/stretchr/testify/mock"
)

type IterxMock struct {
	mock.Mock
}

func (mock IterxMock) Unsafe() igocqlx.IIterx {
	args := mock.Called()

	return args.Get(0).(igocqlx.IIterx)
}

func (mock IterxMock) StructOnly() igocqlx.IIterx {
	args := mock.Called()

	return args.Get(0).(igocqlx.IIterx)
}

func (mock IterxMock) Get(dest interface{}) error {
	args := mock.Called(dest)

	return args.Error(0)
}

func (mock IterxMock) Select(dest interface{}) error {
	args := mock.Called(dest)

	return args.Error(0)
}

func (mock IterxMock) StructScan(dest interface{}) bool {
	args := mock.Called(dest)

	return args.Get(0).(bool)
}

func (mock IterxMock) Scan(dest ...interface{}) bool {
	args := mock.Called(dest)

	return args.Get(0).(bool)
}

func (mock IterxMock) Close() error {
	args := mock.Called()

	return args.Error(0)
}

func (mock IterxMock) MapScan(m map[string]interface{}) bool {
	args := mock.Called(m)

	return args.Bool(0)
}
