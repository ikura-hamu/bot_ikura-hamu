package impl

import (
	"context"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	mongoMigrate "github.com/golang-migrate/migrate/v4/database/mongodb"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/ikura-hamu/bot_ikura-hamu/src/conf"
	"github.com/ikura-hamu/bot_ikura-hamu/src/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

var _ repository.BotRepository = &BotRepository{}

type BotRepository struct {
	db     *mongo.Database
	logger *zap.Logger
}

//go:embed migrate/*.json
var fs embed.FS

func NewBotRepository(l *zap.Logger) *BotRepository {
	mongoDBConfig := conf.GetMongoUri()

	ctx := context.Background()
	option := options.Client().
		SetHosts([]string{mongoDBConfig.Host}).
		SetAuth(options.Credential{Username: mongoDBConfig.User, Password: mongoDBConfig.Password, AuthSource: "admin"})
	c, err := mongo.Connect(ctx, option)
	if err != nil {
		l.Panic("db connection failed", zap.Error(err))
	}

	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		l.Panic("db connection failed", zap.Error(err))
	} else {
		l.Info("connected to db")
	}

	d, err := iofs.New(fs, "migrate")
	if err != nil {
		l.Panic("failed to get io/fs driver", zap.Error(err))
	}

	migrationDriver, err := mongoMigrate.WithInstance(c, &mongoMigrate.Config{DatabaseName: mongoDBConfig.DatabaseName})
	defer func() {
		if err := migrationDriver.Close(); err != nil {
			l.Panic("failed to close migration driver", zap.Error(err))
		}
	}()
	if err != nil {
		l.Panic("failed to create migrate driver", zap.Error(err))
	}

	m, err := migrate.NewWithInstance("iofs", d, "mongodb", migrationDriver)
	if err != nil {
		l.Panic("failed to create migration instance", zap.Error(err))
	}

	current, _, _ := m.Version()
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		l.Panic("failed to migrate", zap.Error(err))
	}
	new, _, _ := m.Version()

	l.Info("migration completed", zap.Uint("previous", current), zap.Uint("new", new))

	db := c.Database(mongoDBConfig.DatabaseName)

	return &BotRepository{
		db:     db,
		logger: l,
	}
}
