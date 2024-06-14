package usecase

import (
	"backendapp/app/domain"
	"context"
)

type IArticleUC interface {
	GetAll(ctx context.Context, page int, pageSize int) (resp []domain.Article, err error)
	GetRow(ctx context.Context, id int) (resp domain.Article, err error)
	Create(ctx context.Context, payload domain.Article) (err error)
	Update(ctx context.Context, payload domain.Article) (err error)
	Delete(ctx context.Context, id int) (err error)
}
