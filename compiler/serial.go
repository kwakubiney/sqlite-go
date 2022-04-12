package compiler

import (
	"encoding/gob"
	"fmt"
	"log"
)

var (
	encoder = gob.NewEncoder(&RowsTableBuffer)
	decoder = gob.NewDecoder(&RowsTableBuffer)
)

func SerializeRow(r Row) {
	err := encoder.Encode(r)
	if err != nil {
		log.Println("encode error:", err)
	}

}

func DeserializeRow() {

	var rows Row

	err := decoder.Decode(&rows)

	for err == nil {
		if err != nil {
			log.Fatal("decode error:", err)
		}
		fmt.Printf("%d %s %s\n", rows.ID, rows.Username, rows.Email)
		err = decoder.Decode(&rows)

	}

}
