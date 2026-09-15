package config

import "os"

type Config struct {
	Address  string
	MySQLDSN string
}

func Load() Config {
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = ":8080"
	}

	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:@tcp(127.0.0.1:3306)/hr_app?parseTime=true"
	}

	return Config{Address: address, MySQLDSN: dsn}
}
