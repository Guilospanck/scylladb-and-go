package gocqlxmock

import (
	"context"

	"github.com/Guilospanck/igocqlx"
	"github.com/stretchr/testify/mock"
)

type SessionxMock struct {
	mock.Mock
}

func (mock SessionxMock) ContextQuery(ctx context.Context, stmt string, names []string) igocqlx.IQueryx {
	args := mock.Called(ctx, stmt, names)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock SessionxMock) Query(stmt string, names []string) igocqlx.IQueryx {
	args := mock.Called(stmt, names)

	return args.Get(0).(igocqlx.IQueryx)
}

func (mock SessionxMock) ExecStmt(stmt string) error {
	args := mock.Called(stmt)

	return args.Error(0)
}

func (mock SessionxMock) AwaitSchemaAgreement(ctx context.Context) error {
	args := mock.Called(ctx)

	return args.Error(0)
}

func (mock SessionxMock) Close() {
	mock.Called()
}
