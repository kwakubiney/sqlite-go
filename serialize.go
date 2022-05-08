package sqlitego

import (
	"encoding/binary"
	"io"
	"log"
	"strings"
)

type Row struct {
	ID       string
	Username string
	Email    string
}

func SerializeRow(r Row, db *DB) error {
	var hdr [4]byte
	arrayOfRowValues := make([]string, 3)
	arrayOfRowValues[0], arrayOfRowValues[1], arrayOfRowValues[2] = string(r.ID), r.Username, r.Email
	stringOfRowValues := strings.Join(arrayOfRowValues, ":")
	binary.BigEndian.PutUint32(hdr[:], uint32(len(stringOfRowValues)))
	WriteToIndexMap(db, r)
	_, err := db.File.Write(hdr[:])
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = io.WriteString(db.File, stringOfRowValues)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
