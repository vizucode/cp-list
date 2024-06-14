package articles

import (
	"backendapp/app/domain"
	"backendapp/app/models"
	"context"
)

func (uc *article) Create(ctx context.Context, payload domain.Article) (err error) {

	modelPayload := models.Article{
		Title:   payload.Title,
		Content: payload.Content,
		Status:  payload.Status,
	}

	err = uc.db.CreateArticle(ctx, modelPayload)
	if err != nil {
		return err
	}

	return nil
}
