package config

import "io/ioutil"

// Config contains a database configuration
type Config struct {
	DB *DBConfig
}

// DBConfig contains the configuration of the database
type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

// GetConfig gets the configuration of the database
func GetConfig() *Config {
	data, err := ioutil.ReadFile("password.txt")
	if err != nil {
		panic(err)
	}

	password := string(data)

	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "guest",
			Password: password,
			Name:     "iot-air-quality",
			Charset:  "utf8",
		},
	}
}
