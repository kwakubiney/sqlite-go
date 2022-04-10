package compiler

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func SerializeRow(row *Row) {
	log.Println("Encoding...")

	var inMemory bytes.Buffer

	encoder := gob.NewEncoder(&inMemory)
	err := encoder.Encode(row)
	if err != nil {
		log.Println("encode error:", err)
	}

}

func DeserializeRow(row *Row) {

	var inMemory bytes.Buffer

	decoder := gob.NewDecoder(&inMemory)
	err := decoder.Decode(&row)
	if err != nil {
		log.Println("decode error:", err)
	}
	fmt.Printf("%d %q %q\n", row.id, row.username, row.email)
}
