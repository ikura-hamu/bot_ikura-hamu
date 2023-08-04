package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	uri := "mongodb://root:password@db:27017/bot?authSource=admin"
	m, err := migrate.New("file://migrate", uri)
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
	m.Up()

	ctx := context.Background()
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("connection error:", err)
	} else {
		fmt.Println("connection success")
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error { return c.JSON(http.StatusOK, "ok!") })
	e.Start(":8080")
}
