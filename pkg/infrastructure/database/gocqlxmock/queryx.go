package gocqlxmock

import (
	"context"

	"github.com/Guilospanck/igocqlx"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/stretchr/testify/mock"
)

type QueryxMock struct {
	mock.Mock
	Ctx   context.Context
	Stmt  string
	Names []string
}

func (mock QueryxMock) WithBindTransformer(tr gocqlx.Transformer) igocqlx.IQueryx {
	args := mock.Called(tr)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) BindStruct(arg interface{}) igocqlx.IQueryx {
	args := mock.Called(arg)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) BindStructMap(arg0 interface{}, arg1 map[string]interface{}) igocqlx.IQueryx {
	args := mock.Called(arg0, arg1)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) BindMap(arg map[string]interface{}) igocqlx.IQueryx {
	args := mock.Called(arg)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) Bind(v ...interface{}) igocqlx.IQueryx {
	args := mock.Called(v)

	return args.Get(0).(igocqlx.IQueryx)
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

func (mock QueryxMock) Iter() igocqlx.IIterx {
	args := mock.Called()

	return args.Get(0).(igocqlx.IIterx)
}

func (mock QueryxMock) Consistency(c gocql.Consistency) igocqlx.IQueryx {
	args := mock.Called(c)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) CustomPayload(customPayload map[string][]byte) igocqlx.IQueryx {
	args := mock.Called(customPayload)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) Trace(trace gocql.Tracer) igocqlx.IQueryx {
	args := mock.Called(trace)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) Observer(observer gocql.QueryObserver) igocqlx.IQueryx {
	args := mock.Called(observer)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) PageSize(n int) igocqlx.IQueryx {
	args := mock.Called(n)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) DefaultTimestamp(enable bool) igocqlx.IQueryx {
	args := mock.Called(enable)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) WithTimestamp(timestamp int64) igocqlx.IQueryx {
	args := mock.Called(timestamp)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) RoutingKey(routingKey []byte) igocqlx.IQueryx {
	args := mock.Called(routingKey)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) WithContext(ctx context.Context) igocqlx.IQueryx {
	args := mock.Called(ctx)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) Prefetch(p float64) igocqlx.IQueryx {
	args := mock.Called(p)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) RetryPolicy(r gocql.RetryPolicy) igocqlx.IQueryx {
	args := mock.Called(r)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) SetSpeculativeExecutionPolicy(sp gocql.SpeculativeExecutionPolicy) igocqlx.IQueryx {
	args := mock.Called(sp)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) Idempotent(value bool) igocqlx.IQueryx {
	args := mock.Called(value)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) SerialConsistency(cons gocql.SerialConsistency) igocqlx.IQueryx {
	args := mock.Called(cons)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) PageState(state []byte) igocqlx.IQueryx {
	args := mock.Called(state)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) NoSkipMetadata() igocqlx.IQueryx {
	args := mock.Called()

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock QueryxMock) Release() {
	mock.Called()
}

func (mock QueryxMock) Scan(dest ...interface{}) error {
	args := mock.Called(dest...)

	return args.Error(0)
}
