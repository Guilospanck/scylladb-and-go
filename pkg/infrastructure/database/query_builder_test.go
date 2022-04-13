package database

import (
	"base/pkg/application/interfaces"
	"base/pkg/infrastructure/database/entities"
	"base/pkg/infrastructure/database/gocqlxmock"
	database_interfaces "base/pkg/infrastructure/database/interfaces"
	"base/pkg/infrastructure/database/models"
	"base/pkg/infrastructure/logger"
	"context"
	"testing"

	"github.com/Guilospanck/igocqlx/table"
)

type QueryBuilderSut struct {
	queryBuilder *queryBuilder[any]

	model   *table.Table
	session database_interfaces.ISession
	queryx  *gocqlxmock.QueryxMock
	logger  interfaces.ILogger

	stmt  string
	names []string
	ctx   context.Context
}

func makeQueryBuilderSut() QueryBuilderSut {
	trackingModel := models.NewTrackingDataTable().Table
	loggerSpy := logger.LoggerSpy{}
	sessionMock := &gocqlxmock.SessionxMock{}
	var session database_interfaces.ISession = sessionMock

	queryBuilder := NewQueryBuider[entities.TrackingDataEntity](
		trackingModel, session, loggerSpy,
	)

	stmt := `INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 08:05+0000','New York',1.0,3.0,17)`
	names := []string{"test"}

	ctx := context.Background()

	queryMock := &gocqlxmock.QueryxMock{
		Ctx:   ctx,
		Stmt:  stmt,
		Names: names,
	}

	return QueryBuilderSut{
		model:   trackingModel,
		logger:  loggerSpy,
		session: sessionMock,
		queryx:  queryMock,

		stmt:  stmt,
		names: names,
		ctx:   ctx,
	}
}

func Test_Insert(t *testing.T) {
	t.Run("", func(t *testing.T) {
		// arrange
		sut := makeQueryBuilderSut()
		sut.session.On("Query", sut.stmt, sut.names).Return(sut.queryx)
		sut.queryx.On("WithContext", context.Background()).Return(sut.queryx)

		// act

		// assert

	})
}
