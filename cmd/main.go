package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/sqlite-go"
)

func main(){
	inputBuffer := sqlitego.NewInputBuffer()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		PrintPrompt()
		scanner.Scan()
		command := scanner.Text()
		inputBuffer.Buffer = command
		var statement sqlitego.Statement

		if strings.HasPrefix(inputBuffer.Buffer, ".") {
			switch sqlitego.DoMetaCommand(inputBuffer) {
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

		sqlitego.ExecuteStatement(statement)
		fmt.Println("Executed")

	}
}

func PrintPrompt() {
	fmt.Printf("db > ")
}
