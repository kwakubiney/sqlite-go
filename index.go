package sqlitego

import (
	"encoding/gob"
	"log"
	"os"
)

func WriteToIndexMap(db *DB, r Row) {
	fileInfo, err := db.File.Stat()
	if err != nil {
		log.Println(err)
	}
	fileOffset := fileInfo.Size()
	db.Bucket[r.ID] = fileOffset
}

func WriteToIndexFile(db *DB) {
	encoder := gob.NewEncoder(db.IndexFile)
	encoder.Encode(db.Bucket)
}

func ReadMapFromIndexFile(db *DB) {
	decoder := gob.NewDecoder(db.IndexFile)
	decoder.Decode(&db.Bucket)
}

func RemoveIndexFile(db *DB) {
	if err := os.Truncate(db.IndexFilePath, 0); err != nil {
		log.Printf("failed to truncate: %v", err)
	}
}
