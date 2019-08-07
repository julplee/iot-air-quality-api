# Go API for IOT Air Quality

see: https://github.com/julplee/iot-air-quality

A RESTful API to store IOT metrics with Go using gorilla/mux (API library) and Gorm (ORM for Go)

## Installation & run

### Download this project

`go get github.com/julplee/iot-air-quality-api`

Before running API server, you must configure the database with your own information in config.go

`
func GetConfig() *Config {
    return &Config{
        DB: &DBConfig{
            Dialect: "mysql",
            Username: "guest",
            Password: "guestpwd",
            Name: "iot-air-quality-db",
            Charset: "utf8"
        }
    }
}
`

### Build & run

`cd iot-air-quality-api`
`go build`
`./iot-air-quality-api`
