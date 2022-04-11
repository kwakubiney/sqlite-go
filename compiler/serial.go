package compiler

import (
	"encoding/gob"
	"fmt"
	"log"
)

func SerializeRow(r []Row) {
	log.Println(r)
	encoder := gob.NewEncoder(&RowsTableBuffer)
	err := encoder.Encode(r)
	log.Println(RowsTableBuffer)
	if err != nil {
		log.Println("encode error:", err)
	}

}

func DeserializeRow() {
	var rowsBuffer = RowsTableBuffer
	rowsTable := make([]Row, 0)
	decoder := gob.NewDecoder(&rowsBuffer)
	err := decoder.Decode(&rowsTable)
	if err != nil {
		log.Println("decode error:", err)
	}
	fmt.Println(rowsTable)
}