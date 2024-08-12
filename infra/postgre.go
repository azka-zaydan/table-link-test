package infra

import (
	"fmt"

	 "github.com/azka-zaydan/table-link-test/configs"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

const (
	maxIdleConnecion  = 10
	maxOpenConnection = 10
)

type PostgreConn struct {
	Read  *sqlx.DB
	Write *sqlx.DB
}

func ProvidePostgreConn(config *configs.Config) *PostgreConn {
	return &PostgreConn{
		Read:  CreatePostgreReadConn(*config),
		Write: CreatePostgreWriteConn(*config),
	}
}

func CreatePostgreReadConn(config configs.Config) *sqlx.DB {
	return CreateDBConnection("read", config.DB.Postgre.Read.Name, config.DB.Postgre.Read.Username, config.DB.Postgre.Read.Password, config.DB.Postgre.Read.Host, config.DB.Postgre.Read.Port)
}

func CreatePostgreWriteConn(config configs.Config) *sqlx.DB {
	return CreateDBConnection("write", config.DB.Postgre.Write.Name, config.DB.Postgre.Write.Username, config.DB.Postgre.Write.Password, config.DB.Postgre.Write.Host, config.DB.Postgre.Write.Port)
}

func CreateDBConnection(name, dbName, username, password, host, port string) *sqlx.DB {
	dc := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s", username, dbName, password, host, port)
	db, err := sqlx.Connect("postgres", dc)
	if err != nil {
		log.
			Fatal().
			Err(err).
			Str("name", name).
			Str("dbName", dbName).
			Str("host", host).
			Str("port", port).
			Msg("Failed connecting to database")
	} else {
		log.
			Info().
			Str("name", name).
			Str("dbName", dbName).
			Str("host", host).
			Str("port", port).
			Msg("Connected to database")
	}

	db.SetMaxIdleConns(maxIdleConnecion)
	db.SetMaxOpenConns(maxOpenConnection)
	return db
}
