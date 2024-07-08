package conf

import (
	"os"
	"strconv"
	// "github.com/joho/godotenv"
)

// config for the backend
func DefaultConfig() *Config {
	return &Config{
		//MySQL: &MySQL{
		//	Host:     "127.0.0.1",
		//	Port:     3306,
		//	DB:       "myblog",
		//	Username: "root",
		//	Password: "12345678",
		//},

		App: &App{
			HttpHost: "127.0.0.1",
			HttpPort: 7082,
		},
	}
}

func ConfigFromEnv() *Config {
	// pwd, _ := os.Getwd()
	// fmt.Println(pwd)
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//dbPortNumber, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	httpPortNumber, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	//fmt.Println(os.Getenv("MYSQL_HOST"))
	//fmt.Println(httpPortNumber)
	return &Config{
		//MySQL: &MySQL{
		//	Host:     os.Getenv("MYSQL_HOST"),
		//	Port:     dbPortNumber,
		//	DB:       os.Getenv("MYSQL_DB"),
		//	Username: os.Getenv("MYSQL_USERNAME"),
		//	Password: os.Getenv("MYSQL_PASSWORD"),
		//},
		App: &App{
			HttpHost: os.Getenv("HTTP_HOST"),
			HttpPort: int64(httpPortNumber),
		},
	}
}
