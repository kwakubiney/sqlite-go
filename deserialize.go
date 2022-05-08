package sqlitego

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
)

func NewDecoder(db *DB) *gob.Decoder {
	return gob.NewDecoder(db.IndexFile)
}

func DeserializeAllRows(db *DB) error{
	for k := range db.Bucket {
		err := DeserializeSpecificRow(db, k)
		if err != nil {
			return err
		}
	}
	return nil	
}

func ParseDecodedRow(row string){
	sliceOfDecodedRowValues := strings.Split(row, ":")
	fmt.Printf("ID: %s, Name: %s, Email: %s\n", sliceOfDecodedRowValues[0], sliceOfDecodedRowValues[1], sliceOfDecodedRowValues[2])
}

func DeserializeSpecificRow(db *DB, id string) error {
	offsetOfRow, ok := db.Bucket[id]
	if !ok {
		return fmt.Errorf("no such row id: %s in the database", id)
	}
	var rowLength [4]byte
	_, err := db.File.ReadAt(rowLength[:], offsetOfRow)
	if err != nil {
		log.Println(err)
		return err
	}
	rowLengthBuffer := bytes.NewReader(rowLength[:])
	var rowLengthAsUInt32 uint32
	err = binary.Read(rowLengthBuffer, binary.BigEndian, &rowLengthAsUInt32)
	if err != nil {
		log.Println(err)
		return err
	}

	rowBuffer := make([]byte, uint8(rowLengthAsUInt32))
	offsetOfRowData := offsetOfRow + 4
	_, err = db.File.ReadAt(rowBuffer[:], offsetOfRowData)
	if err != nil {
		log.Println(err)
		return err
	}	
	ParseDecodedRow(string(rowBuffer))
	return nil
}

