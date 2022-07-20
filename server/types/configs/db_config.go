package configs

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	Host       string `json:"host"`
	DriverName string `json:"driver_name"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	DBName     string `json:"db_name"`
	DBPort     int    `json:"db_port"`
	SslMode    string `json:"ssl_mode"`
	MaxIdle    int    `json:"max_idle"`
	MaxOpen    int    `json:"max_open"`
}

func (conf *DBConfig) GetDBEnvInfo() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", conf.Host, conf.UserName, conf.Password, conf.DBName, conf.DBPort, conf.SslMode)
}

func (conf *DBConfig) GetDBConnection() (*sql.DB, error) {
	database, err := sql.Open(conf.DriverName, conf.GetDBEnvInfo())
	if err != nil {
		return nil, err
	}

	database.SetMaxIdleConns(conf.MaxIdle)
	database.SetMaxOpenConns(conf.MaxOpen)

	return database, nil
}
