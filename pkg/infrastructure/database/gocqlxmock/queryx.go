package gocqlxmock

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/mock"
)

type Transformer func(name string, val interface{}) interface{}

type IQueryx interface {
	WithBindTransformer(tr Transformer) IQueryx
	WithContext(ctx context.Context) IQueryx
	BindStruct(arg interface{}) IQueryx
	BindStructMap(arg0 interface{}, arg1 map[string]interface{}) IQueryx
	bindStructArgs(arg0 interface{}, arg1 map[string]interface{}) ([]interface{}, error)
	BindMap(arg map[string]interface{}) IQueryx
	bindMapArgs(arg map[string]interface{}) ([]interface{}, error)
	Bind(v ...interface{}) IQueryx
	Err() error
	Exec() error
	ExecRelease() error
	ExecCAS() (applied bool, err error)
	ExecCASRelease() (bool, error)
	Get(dest interface{}) error
	GetRelease(dest interface{}) error
	GetCAS(dest interface{}) (applied bool, err error)
	GetCASRelease(dest interface{}) (bool, error)
	Select(dest interface{}) error
	SelectRelease(dest interface{}) error
	Iter() IIterx
}

type QueryxMock struct {
	mock.Mock
	Ctx   context.Context
	Stmt  string
	Names []string
}

func (mock QueryxMock) WithBindTransformer(tr Transformer) IQueryx {
	args := mock.Called(tr)

	return args.Get(0).(IQueryx)
}

func (mock QueryxMock) BindStruct(arg interface{}) IQueryx {
	args := mock.Called(arg)

	return args.Get(0).(IQueryx)
}

func (mock QueryxMock) BindStructMap(arg0 interface{}, arg1 map[string]interface{}) IQueryx {
	args := mock.Called(arg0, arg1)

	return args.Get(0).(IQueryx)
}

func (mock QueryxMock) bindStructArgs(arg0 interface{}, arg1 map[string]interface{}) ([]interface{}, error) {
	args := mock.Called(arg0, arg1)

	return args.Get(0).([]interface{}), args.Error(1)
}

func (mock QueryxMock) BindMap(arg map[string]interface{}) IQueryx {
	args := mock.Called(arg)

	return args.Get(0).(IQueryx)
}

func (mock QueryxMock) bindMapArgs(arg map[string]interface{}) ([]interface{}, error) {
	args := mock.Called(arg)

	return args.Get(0).([]interface{}), args.Error(1)
}

func (mock QueryxMock) Bind(v ...interface{}) IQueryx {
	args := mock.Called(v)

	return args.Get(0).(IQueryx)
}

func (mock QueryxMock) Err() error {
	args := mock.Called()

	return args.Error(0)
}

func (mock QueryxMock) Exec() error {
	args := mock.Called()

	return args.Error(0)
}

func (mock QueryxMock) ExecRelease() error {
	args := mock.Called()

	return args.Error(0)
}

func (mock QueryxMock) ExecCAS() (applied bool, err error) {
	args := mock.Called()

	return args.Get(0).(bool), args.Error(1)
}

func (mock QueryxMock) ExecCASRelease() (bool, error) {
	args := mock.Called()

	return args.Get(0).(bool), args.Error(1)
}

func (mock QueryxMock) Get(dest interface{}) error {
	args := mock.Called(dest)

	return args.Error(0)
}

func (mock QueryxMock) GetRelease(dest interface{}) error {
	args := mock.Called(dest)

	return args.Error(0)
}

func (mock QueryxMock) GetCAS(dest interface{}) (applied bool, err error) {
	args := mock.Called(dest)

	return args.Get(0).(bool), args.Error(1)
}

func (mock QueryxMock) GetCASRelease(dest interface{}) (bool, error) {
	args := mock.Called(dest)

	return args.Get(0).(bool), args.Error(1)
}

func (mock QueryxMock) Select(dest interface{}) error {
	args := mock.Called(dest)

	return args.Error(0)
}

func (mock QueryxMock) SelectRelease(dest interface{}) error {
	args := mock.Called(dest)

	return args.Error(0)
}

func (mock QueryxMock) Iter() IIterx {
	args := mock.Called()

	return args.Get(0).(IIterx)
}

func (mock QueryxMock) WithContext(ctx context.Context) IQueryx {
	args := mock.Called(ctx)

	return args.Get(0).(IQueryx)
}

// "Interface assertion"
var (
	_ IQueryx = QueryxMock{}
	_ IQueryx = queryx{}
)

type queryx struct {
	q *gocql.Query
}

func (q queryx) WithBindTransformer(tr Transformer) IQueryx {
	return q
}

func (q queryx) BindStruct(arg interface{}) IQueryx {
	return q
}

func (q queryx) BindStructMap(arg0 interface{}, arg1 map[string]interface{}) IQueryx {
	return q
}

func (q queryx) bindStructArgs(arg0 interface{}, arg1 map[string]interface{}) ([]interface{}, error) {
	return nil, nil
}

func (q queryx) BindMap(arg map[string]interface{}) IQueryx {
	return q
}

func (q queryx) bindMapArgs(arg map[string]interface{}) ([]interface{}, error) {
	return nil, nil
}

func (q queryx) Bind(v ...interface{}) IQueryx {
	return q
}

func (q queryx) Err() error {
	return nil
}

func (q queryx) Exec() error {
	return nil
}

func (q queryx) ExecRelease() error {
	return nil
}

func (q queryx) ExecCAS() (applied bool, err error) {
	return true, nil
}

func (q queryx) ExecCASRelease() (bool, error) {
	return true, nil
}

func (q queryx) Get(dest interface{}) error {
	return nil
}

func (q queryx) GetRelease(dest interface{}) error {
	return nil
}

func (q queryx) GetCAS(dest interface{}) (applied bool, err error) {
	return true, nil
}

func (q queryx) GetCASRelease(dest interface{}) (bool, error) {
	return true, nil
}

func (q queryx) Select(dest interface{}) error {
	return nil
}

func (q queryx) SelectRelease(dest interface{}) error {
	return nil
}

func (q queryx) Iter() IIterx {
	return iterx{
		i: q.q.Iter(),
	}
}

func (q queryx) WithContext(ctx context.Context) IQueryx {
	return q
}
