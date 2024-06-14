package articles

import "backendapp/app/repositories"

type article struct {
	db repositories.IDatabaseRepo
}

func NewArticle(db repositories.IDatabaseRepo) *article {
	return &article{db: db}
}
