package handlers

import sqlitego "github.com/sqlite-go"

type Handler struct {
	DB *sqlitego.DB
}

func New(db *sqlitego.DB) *Handler {
	return &Handler{
		DB: db,
	}
}
