package main

import (
	"accounts-api/cmd/api/routes"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "accounts-api/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Entrega final - Equipo 5
// @version         1.0
// @description     API para el manejo de cuentas de un usuario
// @termsOfService  http://swagger.io/terms/

// @contact.name   Equipo 5
// @contact.url    http://www.equipo5.io/support
// @contact.email  digitalmoneyhouse.grupo5@gmail.com

// @host      localhost:8083
// @BasePath  /api/v1
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

	r := gin.Default()
	router := routes.NewRouter(r, db)
	router.MapRoutes()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	
	r.GET("/api/v1/accounts/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if err := router.Start(fmt.Sprintf(":%v", port)); err != nil {
		log.Fatal(err)
	}
}
