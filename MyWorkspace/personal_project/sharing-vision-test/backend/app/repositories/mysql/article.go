package mysql

import (
	"backendapp/app/models"
	customerror "backendapp/config/custom_error"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

func (repo *mysql) GetAllArticles(ctx context.Context, page int, pageSize int) (resp []models.Article, err error) {
	offset := (page - 1) * pageSize

	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM articles LIMIT $1 OFFSET $2", page, offset)
	if err != nil {
		log.Println(err)
		return []models.Article{}, customerror.NewCustomError("internal server error", 500)
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Article
		if err = rows.Scan(a.Id, &a.Title, &a.Content, &a.Category, &a.Status, &a.UpdatedAt, &a.CreatedAt); err != nil {
			log.Println(err)
			return resp, customerror.NewCustomError("internal server error", 500)
		}
		resp = append(resp, a)
	}

	if err = rows.Err(); err != nil {
		return []models.Article{}, customerror.NewCustomError("internal server error", 500)
	}

	return resp, nil
}

func (repo *mysql) GetRowArticle(ctx context.Context, id int) (resp models.Article, err error) {

	row := repo.db.QueryRowContext(ctx, "SELECT * FROM articles WHERE id = $1", id)

	if err = row.Scan(&resp.Id, &resp.Title, &resp.Content, &resp.Category, &resp.Status, &resp.UpdatedAt, &resp.CreatedAt); err != nil {
		log.Println(err)
		return resp, customerror.NewCustomError("internal server error", 500)
	}

	return
}

func (repo *mysql) CreateArticle(ctx context.Context, payload models.Article) (err error) {

	_, err = repo.db.ExecContext(ctx, "INSERT INTO articles (title, content, category, status) VALUES ($1, $2, $3, $4)", payload.Title, payload.Content, payload.Category, payload.Status)

	if err != nil {
		log.Println(err)
		return customerror.NewCustomError("internal server error", 500)
	}

	return
}

func (repo *mysql) UpdateArticle(ctx context.Context, selectedColumns []string, payload models.Article) (err error) {
	// Build the SET clause of the update statement
	setClause := make([]string, len(selectedColumns))
	args := make([]interface{}, len(selectedColumns)+1)

	for i, col := range selectedColumns {
		setClause[i] = fmt.Sprintf("%s = $%d", col, i+1)
		switch col {
		case "title":
			args[i] = payload.Title
		case "content":
			args[i] = payload.Content
		case "category":
			args[i] = payload.Category
		case "status":
			args[i] = payload.Status
		default:
			return fmt.Errorf("unknown column: %s", col)
		}
	}
	args[len(selectedColumns)] = payload.Id

	setClauseStr := strings.Join(setClause, ", ")

	query := fmt.Sprintf("UPDATE articles SET %s WHERE id = $%d", setClauseStr, payload.Id)

	_, err = repo.db.ExecContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customerror.NewCustomError("article not found", 404)
		}
		log.Println(err)
		return customerror.NewCustomError("internal server error", 500)
	}

	return nil
}

func (repo *mysql) DeleteArticle(ctx context.Context, id int) (err error) {

	_, err = repo.db.ExecContext(ctx, "DELETE FROM articles WHERE id = $1", id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customerror.NewCustomError("article not found", 404)
		}
		log.Println(err)
		return customerror.NewCustomError("internal server error", 500)
	}

	return
}
