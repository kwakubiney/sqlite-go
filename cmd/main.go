package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/sqlite-go"
	"github.com/sqlite-go/internal/engine"
)

func PrintPrompt() {
	fmt.Printf("db > ")
}

func main() {
	db, err := engine.DbOpen("../db", "../index", 0644, "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
		inputBuffer := sqlitego.NewInputBuffer()
		scanner := bufio.NewScanner(os.Stdin)
		for {
			PrintPrompt()
			scanner.Scan()
			command := scanner.Text()
			inputBuffer.Buffer = command
			var statement engine.Statement
			err := engine.ParseAndExecuteStatement(inputBuffer, db, statement)
			if err != nil {
				log.Println(err)
			}
		}
	}
