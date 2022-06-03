package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/sqlite-go"
	"log"
	"os"
)

func PrintPrompt() {
	fmt.Printf("db > ")
}

func main() {
	cmd := flag.String("cmd", "", "")
	flag.Parse()
	argument := *cmd
	if argument == "cli" {
		inputBuffer := sqlitego.NewInputBuffer()
		scanner := bufio.NewScanner(os.Stdin)

		db, err := sqlitego.DbOpen("db", "index", 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
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
		return
	} else {
		fmt.Println("Unknown arguments. Refer to docs.")
	}

	
}

