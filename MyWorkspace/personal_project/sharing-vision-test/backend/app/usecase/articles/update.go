package articles

import (
	"backendapp/app/domain"
	"backendapp/app/models"
	"context"
)

func (uc *article) Update(ctx context.Context, payload domain.Article) (err error) {

	modelPayload := models.Article{
		Id:      payload.Id,
		Title:   payload.Title,
		Content: payload.Content,
		Status:  payload.Status,
	}

	selectedColumn := []string{}

	if payload.Title != "" {
		selectedColumn = append(selectedColumn, "title")
	}
	if payload.Content != "" {
		selectedColumn = append(selectedColumn, "content")
	}
	if payload.Category != "" {
		selectedColumn = append(selectedColumn, "category")
	}
	if payload.Status != "" {
		selectedColumn = append(selectedColumn, "status")
	}

	err = uc.db.UpdateArticle(ctx, selectedColumn, modelPayload)
	if err != nil {
		return err
	}

	return nil
}
