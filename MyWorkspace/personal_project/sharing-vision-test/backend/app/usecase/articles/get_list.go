package articles

import (
	"backendapp/app/domain"
	"context"
)

func (uc *article) GetAll(ctx context.Context, page int, pageSize int) (resp []domain.Article, err error) {

	listArticleResult, err := uc.db.GetAllArticles(ctx, page, pageSize)
	if err != nil {
		return []domain.Article{}, err
	}

	for _, article := range listArticleResult {
		resp = append(resp, domain.Article{
			Id:        article.Id,
			Title:     article.Title,
			Content:   article.Content,
			Category:  article.Category,
			Status:    article.Status,
			UpdatedAt: article.UpdatedAt,
			CreatedAt: article.CreatedAt,
		})
	}

	return resp, nil
}
