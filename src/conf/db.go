package conf

import (
	"fmt"
	"net/url"
)

type MongoDBConfig struct {
	Host         string
	DatabaseName string
	User         string
	Password     string
}

func GetMongoUri() *MongoDBConfig {
	dbName := url.QueryEscape(getEnvOrDefault("NS_MONGODB_DATABASE", "bot"))
	hostName := url.PathEscape(getEnvOrDefault("NS_MONGODB_HOSTNAME", "db"))
	password := url.QueryEscape(getEnvOrDefault("NS_MONGODB_PASSWORD", "password"))
	port := getEnvOrDefault("NS_MONGODB_PORT", "27017")
	user := url.QueryEscape(getEnvOrDefault("NS_MONGODB_USER", "root"))

	return &MongoDBConfig{
		Host:         fmt.Sprintf("%s:%s", hostName, port),
		DatabaseName: dbName,
		User:         user,
		Password:     password,
	}
}
