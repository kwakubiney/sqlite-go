package sqlitego

import (
	"encoding/gob"
	"fmt"
	"log"
)

type Row struct {
	ID       int32
	Username string
	Email    string
}



func SerializeRow(r Row, encoder *gob.Encoder, db *DB) {
	err := encoder.Encode(r)
	if err != nil {
		log.Println("encode error:", err)
	}
}

func DeserializeRow(decoder *gob.Decoder, db *DB){
	var rows Row
	db.File.Seek(0, 0)
	err := decoder.Decode(&rows)
	for err == nil {
		if err != nil {
			log.Println("decode error:", err)
		}
		fmt.Printf("%d %s %s\n", rows.ID, rows.Username, rows.Email)
		err = decoder.Decode(&rows)
	}
}
