package main

import (
	"bank-api/cmd/api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	router := routes.NewRouter(r)

	router.MapRoutes()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	port := os.Getenv("PORT")
	if err := r.Run(fmt.Sprintf(":%v", port)); err != nil {
		log.Fatal(err)
	}
}
