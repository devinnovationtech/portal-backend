package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/sirupsen/logrus"
)

type mysqlDocumentArchiveRepository struct {
	Conn *sql.DB
}

// NewMysqlDocumentArchiveRepository will create an object that represent the documentArchiveRepository interface
func NewMysqlDocumentArchiveRepository(Conn *sql.DB) domain.DocumentArchiveRepository {
	return &mysqlDocumentArchiveRepository{Conn}
}

var queryJoinDocArchive = `SELECT d.id, d.title, d.excerpt, d.description, d.source, d.mimetype, d.category, d.status,
	d.created_by, d.created_at, d.updated_at FROM document_archives d 
	LEFT JOIN users u
	ON d.created_by = u.id
	WHERE 1=1`

func (r *mysqlDocumentArchiveRepository) fetchQuery(ctx context.Context, query string, args ...interface{}) (result []domain.DocumentArchive, err error) {
	rows, err := r.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.DocumentArchive, 0)
	for rows.Next() {
		docArc := domain.DocumentArchive{}
		userID := uuid.UUID{}
		err = rows.Scan(
			&docArc.ID,
			&docArc.Title,
			&docArc.Excerpt,
			&docArc.Description,
			&docArc.Source,
			&docArc.Mimetype,
			&docArc.Category,
			&docArc.Status,
			&userID,
			&docArc.CreatedAt,
			&docArc.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		docArc.CreatedBy = domain.User{ID: userID}

		result = append(result, docArc)
	}

	return result, nil
}

func (r *mysqlDocumentArchiveRepository) count(ctx context.Context, query string, args ...interface{}) (total int64, err error) {

	err = r.Conn.QueryRow(query, args...).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (r *mysqlDocumentArchiveRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.DocumentArchive, total int64, err error) {
	binds := make([]interface{}, 0)
	query := filterDocArchiveQuery(params, &binds)

	if params.SortBy != "" {
		query += ` ORDER BY d.` + params.SortBy + ` ` + params.SortOrder
	} else {
		query += ` ORDER BY d.created_at DESC `
	}

	total, _ = r.count(ctx, ` SELECT COUNT(1) FROM document_archives d LEFT JOIN users u ON d.created_by = u.id WHERE 1=1 `+query, binds...)
	query = queryJoinDocArchive + query + ` LIMIT ?,? `
	binds = append(binds, params.Offset, params.PerPage)
	res, err = r.fetchQuery(ctx, query, binds...)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (r *mysqlDocumentArchiveRepository) Store(ctx context.Context, body *domain.DocumentArchiveRequest, createdBy string) (err error) {
	query := `INSERT document_archives SET title=?, description=?, excerpt=?, source=?, mimetype=?, category=?, status=?, created_by=?, created_at=?, updated_at=?`

	stmt, err := r.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(ctx,
		body.Title,
		body.Description,
		helpers.MakeExcerpt(body.Description, 150),
		body.Source,
		body.Mimetype,
		body.Category,
		body.Status,
		createdBy,
		time.Now(),
		time.Now(),
	)
	return
}

func (r *mysqlDocumentArchiveRepository) Delete(ctx context.Context, ID int64) (err error) {
	query := `DELETE FROM document_archives WHERE id=?`

	res, err := r.Conn.ExecContext(ctx, query, ID)

	rowsAffect, err := res.RowsAffected()
	if rowsAffect == 0 {
		err = domain.ErrNotFound
	}
	return
}

func (r *mysqlDocumentArchiveRepository) GetByID(ctx context.Context, ID int64) (res domain.DocumentArchive, err error) {
	query := queryJoinDocArchive + ` AND d.id=?`

	userID := uuid.UUID{}
	err = r.Conn.QueryRowContext(ctx, query, ID).Scan(
		&res.ID,
		&res.Title,
		&res.Excerpt,
		&res.Description,
		&res.Source,
		&res.Mimetype,
		&res.Category,
		&res.Status,
		&userID,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	res.CreatedBy = domain.User{ID: userID}

	if err != nil {
		err = domain.ErrNotFound
	}
	return
}

func (r *mysqlDocumentArchiveRepository) TabStatus(ctx context.Context) (res []domain.TabStatusResponse, err error) {
	query := `SELECT status, COUNT(1) FROM document_archives GROUP BY status ORDER BY status`

	rows, err := r.Conn.QueryContext(ctx, query)
	if err != nil {
		return
	}
	res = make([]domain.TabStatusResponse, 0)
	for rows.Next() {
		t := domain.TabStatusResponse{}
		err = rows.Scan(
			&t.Status,
			&t.Count,
		)

		if err != nil {
			logrus.Error(err)
			return
		}
		res = append(res, t)
	}
	return
}

func (r *mysqlDocumentArchiveRepository) Update(ctx context.Context, body *domain.DocumentArchiveRequest, updatedBy string, ID int64) (err error) {
	query := `UPDATE document_archives SET title=?, description=?, excerpt=?, source=?, mimetype=?, category=?, updated_by=?, updated_at=? WHERE id=?`

	stmt, err := r.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx,
		body.Title,
		body.Description,
		helpers.MakeExcerpt(body.Description, 150),
		body.Source,
		body.Mimetype,
		body.Category,
		updatedBy,
		time.Now(),
		ID,
	)

	rowsAffect, err := res.RowsAffected()
	if rowsAffect == 0 {
		err = domain.ErrNotFound
	}

	return
}
