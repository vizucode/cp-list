package articles

import (
	"context"
)

func (uc *article) Delete(ctx context.Context, id int) (err error) {

	err = uc.db.DeleteArticle(ctx, id)
	if err != nil {
		return
	}

	return
}
