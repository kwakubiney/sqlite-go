package sqlitego

import (
	"fmt"
	"os"
)

type DB struct{
    File *os.File
    Path string
    Opened bool
}

func DbOpen(path string, mode os.FileMode) (*DB, error){
    var db = &DB{Opened : true}
    db.Path = path
    flag := os.O_RDWR
    var err error
    if db.File, err = os.OpenFile(db.Path, flag|os.O_CREATE, mode); err != nil{
        db.Close()
        return nil, err
    }
    if os.IsNotExist(err) {
        fmt.Println("Creating new db file...")
        createdFile, err := os.Create(path)
        if err != nil{
            return nil, err
        }
        db.File = createdFile
    }
    return db, nil
}

func (db *DB) Close() error{
    if !db.Opened {
        return nil
    }

    db.Opened = false
    if err := db.File.Close(); err != nil {
        return fmt.Errorf("failed to close db: %s", err)
    }

    db.Path = ""
    return nil
}


