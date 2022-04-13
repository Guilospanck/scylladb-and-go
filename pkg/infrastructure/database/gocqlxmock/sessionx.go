package gocqlxmock

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/mock"
)

type ISessionx interface {
	ContextQuery(ctx context.Context, stmt string, names []string) IQueryx
	Query(stmt string, names []string) IQueryx
	ExecStmt(stmt string) error
	Close()
}

type SessionxMock struct {
	mock.Mock
}

func (mock SessionxMock) ContextQuery(ctx context.Context, stmt string, names []string) IQueryx {
	args := mock.Called(ctx, stmt, names)

	return args.Get(0).(IQueryx)
}

func (mock SessionxMock) Query(stmt string, names []string) IQueryx {
	args := mock.Called(stmt, names)

	return args.Get(0).(IQueryx)
}

func (mock SessionxMock) ExecStmt(stmt string) error {
	args := mock.Called(stmt)

	return args.Error(0)
}

func (mock SessionxMock) Close() {
	mock.Called()
}

// "Interface assertion"
var (
	_ ISessionx = SessionxMock{}
	_ ISessionx = sessionx{}
)

type sessionx struct {
	s *gocql.Session
}

func (s sessionx) ContextQuery(ctx context.Context, stmt string, names []string) IQueryx {
	return queryx{
		q: s.s.Query(stmt, names).WithContext(ctx),
	}
}

func (s sessionx) Query(stmt string, names []string) IQueryx {
	return queryx{
		q: s.s.Query(stmt, names),
	}
}

func (s sessionx) ExecStmt(stmt string) error {
	return s.s.Query(stmt).Exec()
}

func (s sessionx) Close() {
	s.s.Close()
}
