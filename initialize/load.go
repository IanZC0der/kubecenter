package initialize

import (
	"github.com/IanZC0der/kubecenter/global"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func LoadConfigFromEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	global.CONF.App.HttpHost = os.Getenv("HTTP_HOST")
	httpPortNumber, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	global.CONF.App.HttpPort = int64(httpPortNumber)
	return nil
}
