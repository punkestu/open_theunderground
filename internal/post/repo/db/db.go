package db

import "database/sql"

type PostDB struct {
	conn *sql.DB
}

func NewPostDB(conn *sql.DB) *PostDB {
	return &PostDB{conn: conn}
}
