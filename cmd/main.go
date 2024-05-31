package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rsengaravua/go-crud/pkg/books"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

func initDB(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	initConfig()

	port := viper.GetString("PORT")
	dbUrl := viper.GetString("DB_URL")

	router := gin.Default()
	dbHandler, err := initDB(dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	books.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	router.Run(port)
}
