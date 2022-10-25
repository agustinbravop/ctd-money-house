package main

import (
	"ctd-money-house/cmd/api/routes"
	"ctd-money-house/internal/auth"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	// _ = godotenv.Load()
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/dmh")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	keycloackClient, _ := auth.NewKeycloakClient("", "", "", "")

	router := routes.NewRouter(r, db, keycloackClient)
	router.MapRoutes()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
