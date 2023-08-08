package conf

import "os"

func GetTraqClientConf() (string, bool) {
	return os.LookupEnv("BOT_ACCESS_TOKEN")
}
