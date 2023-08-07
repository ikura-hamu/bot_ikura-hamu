package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ikura-hamu/bot_ikura-hamu/src/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type BotRepository struct {
	c      *mongo.Client
	logger *zap.Logger
}

func NewBotRepository(l *zap.Logger) *BotRepository {
	uri := conf.GetMongoUri()

	m, err := migrate.New("file://migrate", uri)
	if err != nil {
		l.Panic("failed to create migration instance", zap.Error(err))
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		l.Panic("failed to migrate", zap.Error(err))
	}

	ctx := context.Background()
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("connection error:", err)
		l.Panic("db connection failed", zap.Error(err))
	} else {
		l.Info("connected to db")
	}

	return &BotRepository{
		c:      c,
		logger: l,
	}
}
