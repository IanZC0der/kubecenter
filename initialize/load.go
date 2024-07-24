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
	global.CONF.System.Provisioner = os.Getenv("PROVISIONER")
	global.CONF.System.MetricsServerUrl = os.Getenv("METRICS_SERVER_URL")
	global.CONF.System.Prometheus.Phost = os.Getenv("PROMETHEUS_HOST")
	prometheusPortNumber, _ := strconv.Atoi(os.Getenv("PROMETHEUS_PORT"))
	global.CONF.System.Prometheus.Pport = int64(prometheusPortNumber)
	global.CONF.System.Prometheus.Pscheme = os.Getenv("PROMETHEUS_SCHEME")
	return nil
}
