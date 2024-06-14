package articles

import (
	"backendapp/app/domain"
	"context"
)

func (uc *article) GetRow(ctx context.Context, id int) (resp domain.Article, err error) {
	rowArticle, err := uc.db.GetRowArticle(ctx, id)
	if err != nil {
		return domain.Article{}, err
	}

	resp = domain.Article{
		Id:        rowArticle.Id,
		Title:     rowArticle.Title,
		Content:   rowArticle.Content,
		Category:  rowArticle.Category,
		Status:    rowArticle.Status,
		UpdatedAt: rowArticle.UpdatedAt,
		CreatedAt: rowArticle.CreatedAt,
	}

	return resp, nil
}
