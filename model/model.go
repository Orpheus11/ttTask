package model

import (
	"gopkg.in/pg.v3"
)

var (
	db *pg.DB
)

func InitConnect(db_user, db_passwd string) {
	db = pg.Connect(&pg.Options{
		User:     db_user,
		Password: db_passwd,
	})
	createSchema(db)
}
func createSchema(db *pg.DB) {
	queries := []string{
		`CREATE TABLE users (id serial, name text, type text)`,
		`CREATE TABLE relationships (id serial, user_id int, other_user_id int, state text, type text)`,
	}
	for _, q := range queries {
		db.Exec(q)
	}
}
