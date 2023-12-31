package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mysqlRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/feedback/repository/mysql"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestStoreSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	body := domain.Feedback{}
	_ = faker.FakeData(&body)

	query := "INSERT feedback SET rating=? , compliments=? , criticism=?, suggestions=?, sector=?, created_at=?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(body.Rating,
		body.Compliments,
		body.Criticism,
		body.Suggestions,
		body.Sector,
		body.CreatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	e := mysqlRepo.NewMysqlFeedbackRepository(db)
	err = e.Store(context.TODO(), &body)
	assert.NotNil(t, err)
}

func TestStoreFailed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	body := domain.Feedback{
		Sector:    "sector",
		CreatedAt: time.Now(),
	}

	query := "INSERT feedback SET rating=? , compliments=? , criticism=?, suggestions=?, sector=?, created_at=?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(body.Rating,
		body.Compliments,
		body.Criticism,
		body.Suggestions,
		body.Sector,
		body.CreatedAt).WillReturnResult(sqlmock.NewErrorResult(err))

	e := mysqlRepo.NewMysqlFeedbackRepository(db)
	err = e.Store(context.TODO(), &body)
	assert.Error(t, err)
}
