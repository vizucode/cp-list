package rest

import (
	"backendapp/app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type rest struct {
	article usecase.IArticleUC
}

func NewRest(
	article usecase.IArticleUC,
) *rest {
	return &rest{
		article: article,
	}
}

func (r *rest) ResponseJson(ctx *gin.Context, message string, data interface{}, metadata interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":  message,
		"data":     data,
		"metadata": metadata,
	})
}

func (r *rest) Register(server *gin.Engine) {
	// v1 := server.Group("/api/v1")

}
