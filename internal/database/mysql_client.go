package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/wlopezob/go-movie-suggester/internal/logs"
)

type MysqlClient struct {
	*sql.DB
}

const (
	host     = "abul.db.elephantsql.com"
	port     = 5432
	user     = "vuzitaei"
	password = "0Yflq8UdUrVJhwCVIVuh8TzwKk_Z0sZX"
	dbname   = "vuzitaei"
)

func NewMySQLClient() *MysqlClient {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		logs.Error("cannot create mysql client")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		logs.Error("Conection not is ready")
		panic(err)
	} else {
		logs.Info("Conection is ready")
	}
	return &MysqlClient{db}
}
