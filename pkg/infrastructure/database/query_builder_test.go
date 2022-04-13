package database

import (
	mocks "base/__mocks__"
	"base/pkg/infrastructure/database/entities"
	"base/pkg/infrastructure/database/models"
	"base/pkg/infrastructure/logger"
	"context"
	"testing"

	"github.com/Guilospanck/gocqlxmock"

	igocqlxtable "github.com/Guilospanck/igocqlx/table"
	"github.com/stretchr/testify/assert"
)

func Test_Insert(t *testing.T) {
	t.Run("Should insert data and have no error", func(t *testing.T) {
		// arrange
		sut := makeQueryBuilderSut[entities.TrackingDataEntity](QueryType(INSERT))

		sut.session.On("Query", sut.stmt, sut.names).Return(sut.queryx)
		sut.queryx.On("WithContext", context.Background()).Return(sut.queryx)
		sut.queryx.On("BindStruct", &mocks.CompleteDataEntity).Return(sut.queryx)
		sut.queryx.On("ExecRelease").Return(nil)

		// act
		err := sut.queryBuilder.Insert(sut.ctx, &mocks.CompleteDataEntity)

		// assert
		assert.NoError(t, err)
		sut.session.AssertExpectations(t)
		sut.session.AssertNumberOfCalls(t, "Query", 1)
		sut.queryx.AssertNumberOfCalls(t, "WithContext", 1)
		sut.queryx.AssertNumberOfCalls(t, "BindStruct", 1)
		sut.queryx.AssertNumberOfCalls(t, "ExecRelease", 1)
		sut.session.AssertCalled(t, "Query", sut.stmt, sut.names)
		sut.queryx.AssertCalled(t, "WithContext", context.Background())
		sut.queryx.AssertCalled(t, "BindStruct", &mocks.CompleteDataEntity)
		sut.queryx.AssertCalled(t, "ExecRelease")
	})
}

type QueryBuilderSut[T any] struct {
	queryBuilder *queryBuilder[T]

	model   igocqlxtable.ITable
	session *gocqlxmock.SessionxMock
	queryx  *gocqlxmock.QueryxMock
	logger  logger.LoggerSpy

	stmt  string
	names []string
	ctx   context.Context
}

func makeQueryBuilderSut[T any](queryType QueryType) QueryBuilderSut[T] {
	trackingModel := models.NewTrackingDataTable().Table
	loggerSpy := logger.LoggerSpy{}
	sessionMock := &gocqlxmock.SessionxMock{}

	queryBuilder := NewQueryBuider[T](
		trackingModel, sessionMock, loggerSpy,
	)

	stmt, names := queryType.returnStatementAndNames()

	ctx := context.Background()

	queryMock := &gocqlxmock.QueryxMock{
		Ctx:   ctx,
		Stmt:  stmt,
		Names: names,
	}

	return QueryBuilderSut[T]{
		model:   trackingModel,
		logger:  loggerSpy,
		session: sessionMock,
		queryx:  queryMock,

		queryBuilder: queryBuilder,

		stmt:  stmt,
		names: names,
		ctx:   ctx,
	}
}

type QueryType string

const (
	INSERT               QueryType = "insert"
	DELETE_PRIMARY_KEY   QueryType = "delete_pk"
	DELETE_PARTITION_KEY QueryType = "delete_partition_key"
	SELECT_PRIMARY_KEY   QueryType = "get"
	SELECT_PARTITION_KEY QueryType = "select_pk"
	SELECT_ALL           QueryType = "select_all"
)

func (qt QueryType) returnStatementAndNames() (string, []string) {
	switch qt {
	case INSERT:
		{
			stmt := `INSERT INTO tracking_data (first_name,last_name,timestamp,heat,location,speed,telepathy_powers) VALUES (?,?,?,?,?,?,?) `
			names := []string{"first_name", "last_name", "timestamp", "heat", "location", "speed", "telepathy_powers"}
			return stmt, names
		}
	case DELETE_PRIMARY_KEY:
		{
			// TODO: implement
		}
	case DELETE_PARTITION_KEY:
		{
			// TODO: implement
		}
	case SELECT_PRIMARY_KEY:
		{
			// TODO: implement
		}
	case SELECT_PARTITION_KEY:
		{
			// TODO: implement
		}
	case SELECT_ALL:
		{
			// TODO: implement
		}
	}

	stmt := `INSERT INTO tracking_data (first_name,last_name,timestamp,heat,location,speed,telepathy_powers) VALUES (?,?,?,?,?,?,?) `
	names := []string{"first_name", "last_name", "timestamp", "heat", "location", "speed", "telepathy_powers"}
	return stmt, names
}
