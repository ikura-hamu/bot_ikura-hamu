package impl

import (
	"context"
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	mongoMigrate "github.com/golang-migrate/migrate/v4/database/mongodb"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/ikura-hamu/bot_ikura-hamu/src/conf"
	"github.com/ikura-hamu/bot_ikura-hamu/src/repository"
	"github.com/pkg/errors"
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
		ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=admin", mongoDBConfig.User, mongoDBConfig.Password, mongoDBConfig.Host, mongoDBConfig.DatabaseName))
	c, err := mongo.Connect(ctx, option)
	if err != nil {
		l.Panic("db connection failed", zap.Error(err))
	}

	db := c.Database(mongoDBConfig.DatabaseName)

	err = db.Client().Ping(ctx, readpref.Primary())
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

	return &BotRepository{
		db:     db,
		logger: l,
	}
}

func handleError(err error) error {
	return errors.WithStack(err)
}
