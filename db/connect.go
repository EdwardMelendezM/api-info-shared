package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/EdwardMelendezM/api-info-shared/config"
)

func InitClients(cfg config.Configuration) (err error) {
	var (
		MySqlUri = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s&multiStatements=true",
			cfg.DB.DbUsername,
			cfg.DB.DbPassword,
			cfg.DB.DbHost,
			cfg.DB.DbPort,
			cfg.DB.DbDatabase,
			"America%2FLima",
		)
	)
	Client, err = ConnectMySQL(MySqlUri)
	if err != nil {
		return err
	}
	return nil
}

func ConnectMySQL(uri string) (client *sql.DB, err error) {
	client, err = sql.Open("mysql", uri)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MySQL", uri)
	client.SetConnMaxLifetime(time.Minute * 6)
	client.SetMaxOpenConns(200)
	client.SetMaxIdleConns(200)
	return client, nil
}
