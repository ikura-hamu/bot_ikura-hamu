package conf

import "os"

func GetTraqClientConf() (string, bool) {
	return os.LookupEnv("BOT_ACCESS_TOKEN")
}

func GetBotToken() string {
	return os.Getenv("BOT_VERIFICATION_TOKEN")
}

func GetBotUserId() string {
	return os.Getenv("BOT_USER_ID")
}
