package mysql

import "database/sql"

type mysql struct {
	db *sql.DB
}

func NewMysql(db *sql.DB) *mysql {
	return &mysql{
		db: db,
	}
}
