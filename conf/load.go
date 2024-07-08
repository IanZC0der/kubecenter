package conf

//var (
//	config *Config = DefaultConfig()
//	//config *Config = ConfigFromEnv()
//)
//
//func C() *Config {
//	return config
//}

func LoadConfigFromToml(filepath string) error {

	// conf := DefaultConfig()
	// _, err := toml.DecodeFile(filepath, config)

	// if err != nil {
	// 	return err
	// }

	return nil
}

//func LoadConfigFromEnv() error {
//	config.MySQL.Host = os.Getenv("MYSQL_HOST")
//	portNumber, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
//	config.MySQL.Port = portNumber
//	config.MySQL.Username = os.Getenv("MYSQL_USER")
//	config.MySQL.Password = os.Getenv("MYSQL_PASSWORD")
//	config.MySQL.DB = os.Getenv("MYSQL_DB")
//	config.App.HttpHost = os.Getenv("HTTP_HOST")
//	portNumber, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))
//	config.App.HttpPort = int64(portNumber)
//
//	return nil
//}
