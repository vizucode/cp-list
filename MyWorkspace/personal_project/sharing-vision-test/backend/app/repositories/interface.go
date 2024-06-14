package repositories

import (
	"backendapp/app/models"
	"context"
)

type IDatabaseRepo interface {
	GetAllArticles(ctx context.Context, page int, pageSize int) (resp []models.Article, err error)
	GetRowArticle(ctx context.Context, id int) (resp models.Article, err error)
	CreateArticle(ctx context.Context, payload models.Article) (err error)
	UpdateArticle(ctx context.Context, selectedColumn []string, payload models.Article) (err error)
	DeleteArticle(ctx context.Context, id int) (err error)
}
