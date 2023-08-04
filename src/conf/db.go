package conf

import "fmt"

func GetMongoUri() string {
	dbName := getEnvOrDefault("NS_MONGODB_DATABASE", "bot")
	hostName := getEnvOrDefault("NS_MONGODB_HOSTNAME", "db")
	password := getEnvOrDefault("NS_MONGODB_PASSWORD", "password")
	port := getEnvOrDefault("NS_MONGODB_PORT", "27017")
	user := getEnvOrDefault("NS_MONGO_USER", "root")

	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", user, password, hostName, port, dbName)
}
