package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mysqlRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/news/repository/mysql"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockNews := []domain.News{
		{
			ID:        1,
			Title:     "title",
			Excerpt:   "excerpt",
			Content:   "content",
			Views:     10,
			Shared:    25,
			Image:     domain.NullString{String: "image", Valid: true},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "title",
			Excerpt:   "excerpt",
			Content:   "content",
			Views:     15,
			Shared:    30,
			Image:     domain.NullString{String: "image 2", Valid: true},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "category", "title", "excerpt", "content", "image", "video", "slug", "author_id", "type", "views", "shared", "source", "created_at", "updated_at"}).
		AddRow(mockNews[0].ID, mockNews[0].Category, mockNews[0].Title, mockNews[0].Excerpt, mockNews[0].Content,
			nil, nil, mockNews[0].Slug.String, "", "", mockNews[0].Views, mockNews[0].Shared, mockNews[0].Source.String, mockNews[0].CreatedAt, mockNews[0].UpdatedAt).
		AddRow(mockNews[1].ID, mockNews[1].Category, mockNews[1].Title, mockNews[1].Excerpt, mockNews[1].Content,
			nil, nil, mockNews[1].Slug.String, "", "", mockNews[1].Views, mockNews[1].Shared, mockNews[1].Source.String, mockNews[1].CreatedAt, mockNews[1].UpdatedAt)

	query := "SELECT id, category, title, excerpt, content, image, video, slug, author_id, type, views, shared, source, created_at, updated_at FROM news"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := mysqlRepo.NewMysqlNewsRepository(db)

	params := &domain.Request{
		Keyword:   "",
		PerPage:   10,
		Offset:    0,
		SortBy:    "",
		SortOrder: "",
	}

	list, _, err := a.Fetch(context.TODO(), params)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}
