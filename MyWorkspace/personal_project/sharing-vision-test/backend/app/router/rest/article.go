package rest

import (
	"backendapp/app/domain"
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *rest) GetListArticle(c *gin.Context) {
	ctx := context.Background()

	var (
		page     int
		pageSize int
	)

	pageStr, isPage := c.Params.Get("page")
	if isPage {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSizeStr, isPageSize := c.Params.Get("pageSize")

	if isPageSize {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	resp, err := r.article.GetAll(ctx, page, pageSize)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(c, "success get articles", resp, nil)
}

func (r *rest) GetRowArticle(c *gin.Context) {
	ctx := context.Background()

	var (
		id int
	)

	idStr, isExist := c.Params.Get("id")
	if isExist {
		id, _ = strconv.Atoi(idStr)
	}

	resp, err := r.article.GetRow(ctx, id)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(c, "success get article", resp, nil)
}

func (r *rest) CreateArticle(c *gin.Context) {
	ctx := context.Background()

	var (
		payload domain.Article
	)

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	err = r.article.Create(ctx, payload)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(c, "success create article", nil, nil)
}

func (r *rest) UpdateArticle(c *gin.Context) {
	ctx := context.Background()

	var (
		payload domain.Article
		id      int
	)

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	idStr, isExist := c.Params.Get("id")
	if isExist {
		id, _ = strconv.Atoi(idStr)
	}
	payload.Id = id

	err = r.article.Update(ctx, payload)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(c, "success update article", nil, nil)
}

func (r *rest) DeleteArticle(c *gin.Context) {
	ctx := context.Background()

	var (
		id int
	)

	idStr, isExist := c.Params.Get("id")
	if isExist {
		id, _ = strconv.Atoi(idStr)
	}

	err := r.article.Delete(ctx, id)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(c, "success delete article", nil, nil)
}
