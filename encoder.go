package sqlitego

import (
	"encoding/gob"
)

func NewEncoder(db *DB) *gob.Encoder{
	return gob.NewEncoder(db.File)
}

