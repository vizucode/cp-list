package errorhandling

import (
	customerror "backendapp/config/custom_error"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				if e, ok := err.(*customerror.CustomError); ok {
					ctx.AbortWithStatusJSON(e.HttpCode, gin.H{
						"message": e.Message,
					})
				}

				if validationErr, ok := err.(validator.ValidationErrors); ok {
					errsMap := make(map[string]interface{})
					for _, val := range validationErr {
						switch val.Tag() {
						case "required":
							errsMap[val.Field()] = "This field is required"
						case "min":
							errsMap[val.Field()] = "This field must be at least 8 characters long"
						case "email":
							errsMap[val.Field()] = "This field must be a valid email address"
						case "code_number_id":
							errsMap[val.Field()] = "This field must started with 62"
						default:
							errsMap[val.Field()] = "This field is invalid"
						}
					}

					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": errsMap,
					})
				}
			}
		}()
		ctx.Next()
	}
}
