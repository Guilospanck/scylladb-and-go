package database

import (
	"base/pkg/application/interfaces"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
	"go.uber.org/zap"
)

type QueryBuilder[T interface{}] struct {
	Model   *table.Table
	Session *gocqlx.Session
	Logger  interfaces.ILogger
}

func (queryBuilder *QueryBuilder[T]) SelectAll() ([]T, error) {
	selectStatement, statementNames := queryBuilder.Model.SelectAll()
	selectQuery := queryBuilder.Session.Query(selectStatement, statementNames)

	var results []T
	err := selectQuery.SelectRelease(&results)
	if err != nil {
		queryBuilder.Logger.Error("SelectAll-SelectRelease() error", zap.Error(err))
		return nil, err
	}

	return results, nil
}

// func NewQueryBuider(model *table.Table, session *gocqlx.Session, logger interfaces.ILogger) *queryBuilder {
// 	return &queryBuilder{
// 		model,
// 		session,
// 		logger,
// 	}
// }
