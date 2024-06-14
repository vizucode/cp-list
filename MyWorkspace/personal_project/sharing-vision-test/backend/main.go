package main

import (
	errorhandling "backendapp/app/middlewares/error_handling"
	"backendapp/app/repositories/mysql"
	"backendapp/app/router/rest"
	"backendapp/app/usecase/articles"
	"backendapp/config/database"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()
	port := os.Getenv("PORT")
	if strings.EqualFold(port, "") {
		port = ":8080"
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	server.Use(errorhandling.ErrorHandler())
	dbConn := database.GetDB()
	defer dbConn.Close()

	// repository
	mysql := mysql.NewMysql(dbConn)

	// service
	articleUC := articles.NewArticle(mysql)

	rest.NewRest(
		articleUC,
	).Register(server)

	server.Run(port)
}
