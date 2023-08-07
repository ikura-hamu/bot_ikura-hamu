package impl

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ikura-hamu/bot_ikura-hamu/src/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type BotRepository struct {
	c *mongo.Client
}

func NewBotRepository() *BotRepository {
	uri := conf.GetMongoUri()

	m, err := migrate.New("file://migrate", uri)
	if err != nil {
		log.Fatalf("failed to create migration instance: %v", err)
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("failed to migrate: %v", err)
	}

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

	return &BotRepository{c: c}
}
