package conf

import (
	"log"
	"strings"
)

func GetMongoUri() string {
	dbName := getEnvOrDefault("NS_MONGODB_DATABASE", "bot")
	hostName := getEnvOrDefault("NS_MONGODB_HOSTNAME", "db")
	password := getEnvOrDefault("NS_MONGODB_PASSWORD", "password")
	port := getEnvOrDefault("NS_MONGODB_PORT", "27017")
	user := getEnvOrDefault("NS_MONGODB_USER", "root")

	// return fmt.Sprintf("mongodb://%s:%v@%s:%s/%s?authSource=admin", user, password, hostName, port, dbName)
	uri := "mongodb://" + user + ":" + password + "@" + hostName + ":" + port + "/" + dbName + "?authSource=admin"
	uri = strings.ReplaceAll(uri, "%", "%25")
	log.Println(uri)
	return uri
}
