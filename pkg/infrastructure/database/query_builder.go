package database

import (
	"base/pkg/application/interfaces"
	"context"
	"fmt"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
)

type queryBuilder[T any] struct {
	model   *table.Table
	session *gocqlx.Session
	logger  interfaces.ILogger
}

/* It will insert data into table.
INSERT INTO table VALUES {};
*/
func (queryBuilder *queryBuilder[T]) Insert(ctx context.Context, insertData *T) error {
	insertStatement, insertNames := queryBuilder.model.Insert()
	insertQuery := queryBuilder.session.Query(insertStatement, insertNames).WithContext(ctx)

	err := insertQuery.BindStruct(insertData).ExecRelease()
	if err != nil {
		queryBuilder.logger.Error(fmt.Sprintf("Insert() error %s", err.Error()))
		return err
	}

	return nil
}

/* It will delete from table based on the Primary Key (Partition Key + Clustering Key (if exists))
DELETE FROM table WHERE PK = {};
*/
func (queryBuilder *queryBuilder[T]) Delete(ctx context.Context, dataToBeDeleted *T) error {
	deleteStatement, deleteNames := queryBuilder.model.Delete()
	deleteQuery := queryBuilder.session.Query(deleteStatement, deleteNames).WithContext(ctx)

	err := deleteQuery.BindStruct(dataToBeDeleted).WithContext(ctx).ExecRelease()
	if err != nil {
		queryBuilder.logger.Error(fmt.Sprintf("Delete by Primary Key error: %s", err.Error()))
		return err
	}

	return nil
}

func (queryBuilder *queryBuilder[T]) DeleteAllFromPartitioningKey(ctx context.Context, dataToBeDeleted *T) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE ", queryBuilder.model.Metadata().Name)

	for index, value := range queryBuilder.model.Metadata().PartKey {
		if index == 0 {
			query += fmt.Sprintf("%s=? ", value)
			continue
		}

		query += fmt.Sprintf("AND %s=?", value)
	}

	deleteQuery := queryBuilder.session.Query(query, queryBuilder.model.Metadata().PartKey).WithContext(ctx)

	err := deleteQuery.BindStruct(dataToBeDeleted).WithContext(ctx).ExecRelease()
	if err != nil {
		queryBuilder.logger.Error(fmt.Sprintf("Delete by Partition Key error: %s", err.Error()))
		return err
	}

	return nil
}

/* It will return data based on the Partition Key
SELECT * FROM table WHERE {partition key = {}};
*/
func (queryBuilder *queryBuilder[T]) Select(ctx context.Context, dataToGet *T) ([]T, error) {
	selectStatement, selectNames := queryBuilder.model.Select()
	selectQuery := queryBuilder.session.Query(selectStatement, selectNames).WithContext(ctx)

	var results []T
	err := selectQuery.BindStruct(dataToGet).WithContext(ctx).SelectRelease(&results)
	if err != nil {
		queryBuilder.logger.Error(fmt.Sprintf("Select error: %s", err.Error()))
		return nil, err
	}

	return results, nil
}

/* It will return data based on the Primary Key (Partition + Clustering key)
SELECT * FROM table WHERE {primary key = {}};
*/
func (queryBuilder *queryBuilder[T]) Get(ctx context.Context, dataToGet *T) (*T, error) {
	selectStatement, selectNames := queryBuilder.model.Get()
	selectQuery := queryBuilder.session.Query(selectStatement, selectNames).WithContext(ctx)

	var result []T
	err := selectQuery.BindStruct(dataToGet).WithContext(ctx).SelectRelease(&result)
	if err != nil {
		queryBuilder.logger.Error(fmt.Sprintf("Get error: %s", err.Error()))
		return nil, err
	}

	if len(result) > 0 {
		return &result[0], nil
	}

	return nil, nil
}

/* It will everything from table.
SELECT * FROM table;
*/
func (queryBuilder *queryBuilder[T]) SelectAll(ctx context.Context) ([]T, error) {
	selectAllStatement, selectAllNames := queryBuilder.model.SelectAll()
	selectAllQuery := queryBuilder.session.Query(selectAllStatement, selectAllNames).WithContext(ctx)

	var results []T
	err := selectAllQuery.WithContext(ctx).SelectRelease(&results)
	if err != nil {
		queryBuilder.logger.Error(fmt.Sprintf("SelectAll error: %s", err.Error()))
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
