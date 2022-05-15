package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sqlite-go"
)

func main() {
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
		fmt.Println(command)
		inputBuffer.Buffer = command
		var statement sqlitego.Statement

		if strings.HasPrefix(inputBuffer.Buffer, ".") {
			switch sqlitego.DoMetaCommand(inputBuffer, db) {
			case sqlitego.MetaCommandSuccess:
				continue
			case sqlitego.MetaCommandUnrecognizedCommand:
				fmt.Printf("Unrecognized command %q \n", inputBuffer.Buffer)
				continue
			}
		}

		switch sqlitego.PrepareStatement(inputBuffer, &statement) {
		case sqlitego.PrepareSuccess:

		case sqlitego.PrepareUnrecognizedStatement:
			fmt.Printf("Unrecognized command at start of %q \n", inputBuffer.Buffer)
			continue

		case sqlitego.PrepareSyntaxError:
			fmt.Println("Syntax error. Could not parse statement.")
			continue
		}

		err := sqlitego.ExecuteStatement(statement, db)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("Executed")

	}

}

func PrintPrompt() {
	fmt.Printf("db > ")
}
