package main

import (
	"ctd-money-house/cmd/api/routes"
	"ctd-money-house/internal/auth"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@(%s)/%s", dbUsername, dbPass, dbHost, dbDatabase))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	kcUrl := os.Getenv("KEYCLOAK_URL")
	kcRealm := os.Getenv("KEYCLOAK_REALM")
	kcClientID := os.Getenv("KEYCLOAK_CLIENT_ID")
	kcClientSecret := os.Getenv("KEYCLOAK_CLIENT_SECRET")
	keycloakClient, err := auth.NewKeycloakClient(kcUrl, kcClientID, kcClientSecret, kcRealm)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	router := routes.NewRouter(r, db, keycloakClient)
	router.MapRoutes()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	port := os.Getenv("PORT")
	if err := r.Run(fmt.Sprintf(":%v", port)); err != nil {
		log.Fatal(err)
	}
}
