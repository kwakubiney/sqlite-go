package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sqlite-go"
	"github.com/sqlite-go/handlers"
	"github.com/sqlite-go/server"
)

func PrintPrompt() {
	fmt.Printf("db > ")
}

func main() {
	cmd := flag.String("cmd", "", "")
	flag.Parse()
	argument := *cmd
	db, err := sqlitego.DbOpen("db", "index", 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if argument == "cli" {
		inputBuffer := sqlitego.NewInputBuffer()
		scanner := bufio.NewScanner(os.Stdin)
		for {
			PrintPrompt()
			scanner.Scan()
			command := scanner.Text()
			inputBuffer.Buffer = command
			var statement sqlitego.Statement
			err := sqlitego.ParseAndExecuteStatement(inputBuffer, db, statement)
			if err != nil {
				log.Fatal(err)
			}
		}

	} else if argument == "http" {
		handlers := handlers.New(db)
		server := server.New(handlers)
		server.Start()
	} else {
		fmt.Println("Unknown arguments. Refer to docs.")
	}
}
