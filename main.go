package main

import (
	"context"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/sh0e1/wire/infrastructure"
)

func main() {
	sqlHandler, err := infrastructure.NewSQLHandler(infrastructure.SQLite, "./data.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlHandler.Close()

	if err := migrate(sqlHandler); err != nil {
		log.Fatal(err)
	}

	route := infrastructure.Route(sqlHandler)
	if err := http.ListenAndServe(":8080", route); err != nil {
		log.Fatal(err)
	}
}

func migrate(sqlHander *infrastructure.SQLHandler) error {
	const sqlstr = `create table if not exists todos (` +
		`id string not null primary key,` +
		`title string` +
		`)`

	_, err := sqlHander.ExecuteContext(context.Background(), sqlstr)
	return err
}
