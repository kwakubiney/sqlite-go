package sqlitego

import "encoding/gob"

func NewDecoder(db *DB) *gob.Decoder{
	return gob.NewDecoder(db.IndexFile)
}
 