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

/* It'll insert data into table.
*  INSERT INTO table VALUES {};
 */
func (queryBuilder *queryBuilder[T]) Insert(insertData *T) error {
	insertStatement, insertNames := queryBuilder.model.Insert()
	insertQuery := queryBuilder.session.Query(insertStatement, insertNames)

	err := insertQuery.BindStruct(insertData).ExecRelease()
	if err != nil {
		queryBuilder.logger.Error("Insert error: ", zap.Error(err))
		return err
	}

	return nil
}

/* It'll delete from table based on the Primary Key (Partition Key + Clustering Key (if exists))
*  DELETE FROM table WHERE PK = {};
 */
func (queryBuilder *queryBuilder[T]) Delete(dataToBeDeleted *T) error {
	deleteStatement, deleteNames := queryBuilder.model.Delete()
	deleteQuery := queryBuilder.session.Query(deleteStatement, deleteNames)

	err := deleteQuery.BindStruct(dataToBeDeleted).ExecRelease()
	if err != nil {
		queryBuilder.logger.Error("Delete error: ", zap.Error(err))
		return err
	}

	return nil
}

/* It'll return data based on the Partition Key
*  SELECT * FROM table WHERE {partition key = {}};
 */
func (queryBuilder *queryBuilder[T]) Select(dataToGet *T) ([]T, error) {
	selectStatement, selectNames := queryBuilder.model.Select()
	selectQuery := queryBuilder.session.Query(selectStatement, selectNames)

	var results []T
	err := selectQuery.BindStruct(dataToGet).SelectRelease(&results)
	if err != nil {
		queryBuilder.logger.Error("Select-BindStruct-SelectRelease() error", zap.Error(err))
		return nil, err
	}

	return results, nil
}

/* It'll everything from table.
*  SELECT * FROM table;
 */
func (queryBuilder *queryBuilder[T]) SelectAll() ([]T, error) {
	selectAllStatement, selectAllNames := queryBuilder.model.SelectAll()
	selectAllQuery := queryBuilder.session.Query(selectAllStatement, selectAllNames)

	var results []T
	err := selectAllQuery.SelectRelease(&results)
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
