package compiler

import (
	"encoding/gob"
	"fmt"
	"log"
)

func SerializeRow(r []Row) {

	encoder := gob.NewEncoder(&RowsTableBuffer)
	err := encoder.Encode(r)
	if err != nil {
		log.Println("encode error:", err)
	}

}

func DeserializeRow() {
	var rowsBuffer = RowsTableBuffer
	rowsTable := make([]Row, 0)
	log.Println(rowsTable)

	decoder := gob.NewDecoder(&rowsBuffer)

	err := decoder.Decode(&rowsTable)
	if err != nil {
		log.Println("encode error:", err)
	}
	log.Println(len(rowsTable))
	fmt.Println(rowsTable)
	// fmt.Printf("%d %q %q\n", row.Id, row.Username, row.Email)
}
