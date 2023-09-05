package conf

import (
	"fmt"
)

type MongoDBConfig struct {
	Host         string
	DatabaseName string
	User         string
	Password     string
}

func GetMongoUri() *MongoDBConfig {
	dbName := getEnvOrDefault("NS_MONGODB_DATABASE", "bot")
	hostName := getEnvOrDefault("NS_MONGODB_HOSTNAME", "db")
	password := getEnvOrDefault("NS_MONGODB_PASSWORD", "password")
	port := getEnvOrDefault("NS_MONGODB_PORT", "27017")
	user := getEnvOrDefault("NS_MONGODB_USER", "root")

	return &MongoDBConfig{
		Host:         fmt.Sprintf("%s:%s", hostName, port),
		DatabaseName: dbName,
		User:         user,
		Password:     password,
	}
}
