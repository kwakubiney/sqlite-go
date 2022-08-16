package engine

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
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	var stringLength [4]byte
	arrayOfRowValues := make([]string, 3)
	arrayOfRowValues[0], arrayOfRowValues[1], arrayOfRowValues[2] = string(r.ID), r.Username, r.Email
	stringOfRowValues := strings.Join(arrayOfRowValues, ":")
	binary.BigEndian.PutUint32(stringLength[:], uint32(len(stringOfRowValues)))
	PushToIndexMapWithoutLock(db, r)
	_, err := db.File.Write(stringLength[:])
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

func SerializeRowToBytes()

