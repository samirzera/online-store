package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBandoDeDados() *sql.DB {
	conexao := "user=postgres dbname=loja_samir password=09102002 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
