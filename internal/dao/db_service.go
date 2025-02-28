package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type sslModes string

const (
	Required   sslModes = "require"
	Disabled   sslModes = "disable"
	VerifyFull sslModes = "verify-full"
	VerifyCa   sslModes = "verify-ca"
)

type DbConfig struct {
	DbName             string
	DbUser             string
	DbPass             string
	DbPort             int
	DbHost             string
	SslMode            *string
	MaxIdleConn        int
	MaxOpenConn        int
	MaxLifetimeMinutes int
	MaxIdleTimeSecs    int
}

const (
	maxIdleConn    = 10
	maxIdleDbConn  = 5
	maxDbLifetime  = 5 * time.Minute
	defaultSslMode = "require"
)

var db *sql.DB

type DbConnection struct {
	DbConn *sql.DB
}

func InitDatabase(dbConfig DbConfig) (DbConnection, error) {
	sslMode := checkSslMode(dbConfig.SslMode)

	dbConnInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbUser, dbConfig.DbPass, dbConfig.DbName, sslMode)

	connection := DbConnection{}

	if db != nil {
		connection.DbConn = db
	} else {
		pDb, err := sql.Open("postgres", dbConnInfo)
		if err != nil {
			return DbConnection{}, err
		}

		connection.DbConn = pDb
	}

	if dbConfig.MaxIdleConn == 0 {
		connection.DbConn.SetMaxOpenConns(maxIdleConn)
	}

	if dbConfig.MaxOpenConn == 0 {
		connection.DbConn.SetMaxOpenConns(maxIdleDbConn)
	} else {
		connection.DbConn.SetMaxOpenConns(dbConfig.MaxOpenConn)
	}

	if dbConfig.MaxLifetimeMinutes != 0 {
		connection.DbConn.SetConnMaxLifetime(time.Minute * time.Duration(dbConfig.MaxLifetimeMinutes))
	}

	if dbConfig.MaxIdleTimeSecs != 0 {
		connection.DbConn.SetConnMaxIdleTime(time.Second * time.Duration(dbConfig.MaxIdleTimeSecs))
	}

	return connection, nil
}

func (d DbConnection) PingDb() error {
	err := d.DbConn.Ping()
	if err != nil {
		return err
	}
	return nil
}

func checkSslMode(sslMode *string) sslModes {

	if sslMode == nil {
		return defaultSslMode
	}

	switch sslModes(*sslMode) {
	case Required:
		return Required
	case Disabled:
		return Disabled
	case VerifyFull:
		return VerifyFull
	case VerifyCa:
		return VerifyCa
	default:
		return defaultSslMode
	}
}
