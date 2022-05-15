package sqlitego

import (
	"fmt"
	"log"
	"os"
)

func (db *DB) TestClose() error {
	if !db.Opened {
		return nil
	}
	db.Opened = false

	//WriteToIndexFile(db)
	RemoveIndexFile(db)
	RemoveDBTestFile()

	db.Path = ""
	db.IndexFilePath = ""

	if err := db.File.Close(); err != nil {
		return fmt.Errorf("failed to close db: %s", err)
	}

	if err := db.IndexFile.Close(); err != nil {
		return fmt.Errorf("failed to close index file: %s", err)
	}

	return nil
}

//TODO: This defeats DRY principle, have to create a general function to truncate files based on file names.
func RemoveDBTestFile() {
	if err := os.Truncate("test-db", 0); err != nil {
		log.Printf("failed to truncate: %v", err)
	}
}
