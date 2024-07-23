package conf

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type App struct {
	HttpHost string `json:"http_host" yaml:"http_host" env:"HTTP_HOST"`
	HttpPort int64  `json:"http_port" yaml:"http_port" env:"HTTP_PORT"`
}

type System struct {
	Provisioner      string `json:"provisioner" yaml:"provisioner" env:"PROVISIONER"`
	MetricsServerUrl string `json:"metrics_server_url" yaml:"metrics_server_url" env:"METRICS_SERVER_URL"`
}

type Config struct {
	//MySQL *MySQL `json:"mysql"`

	App    *App    `json:"app" yaml:"app"`
	System *System `json:"system" yaml:"system"`
}

func (a *App) HttpAddress() string {
	return fmt.Sprintf("%s:%d", a.HttpHost, a.HttpPort)
}

func (c *Config) String() string {

	jsonConfig, _ := json.Marshal(c)
	return string(jsonConfig)
}

type MySQL struct {
	Host     string `json:"host" toml:"host" env:"MYSQL_HOST"`
	Port     int    `json:"port" toml:"port" env:"MYSQL_PORT"`
	DB       string `json:"database" toml:"database" env:"MYSQL_DB"`
	Username string `json:"username" toml:"username" env:"MYSQL_USERNAME"`
	Password string `json:"password" toml:"password" env:"MYSQL_PASSWORD"`

	lock sync.Mutex
	conn *gorm.DB
}

func (m *MySQL) GetDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB)
	return dsn
}

func (m *MySQL) GetConn() *gorm.DB {
	if m.conn == nil {
		// return m.conn
		m.lock.Lock()
		defer m.lock.Unlock()
		conn, err := gorm.Open(mysql.Open(m.GetDSN()), &gorm.Config{})

		if err != nil {
			panic(err)
		}
		m.conn = conn

	} // if db is nil, create a new connection otherwise return the existing connection

	return m.conn
}
