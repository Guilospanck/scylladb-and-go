package database

import (
	"base/pkg/application/interfaces"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
	"go.uber.org/zap"
)

type queryBuilder[T any] struct {
	model   *table.Table
	session *gocqlx.Session
	logger  interfaces.ILogger
}

func (queryBuilder *queryBuilder[T]) SelectAll() ([]T, error) {
	selectStatement, statementNames := queryBuilder.model.SelectAll()
	selectQuery := queryBuilder.session.Query(selectStatement, statementNames)

	var results []T
	err := selectQuery.SelectRelease(&results)
	if err != nil {
		queryBuilder.logger.Error("SelectAll-SelectRelease() error", zap.Error(err))
		return nil, err
	}

	return results, nil
}

func NewQueryBuider[T any](model *table.Table, session *gocqlx.Session, logger interfaces.ILogger) *queryBuilder[T] {
	return &queryBuilder[T]{
		model,
		session,
		logger,
	}
}
