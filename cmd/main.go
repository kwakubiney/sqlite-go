package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sqlite-go/compiler"
)

func main() {
	inputBuffer := compiler.NewInputBuffer()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		PrintPrompt()
		scanner.Scan()
		command := scanner.Text()
		inputBuffer.Buffer = command

		if strings.HasPrefix(inputBuffer.Buffer, ".") {
			switch compiler.DoMetaCommand(inputBuffer) {
			case compiler.MetaCommandSuccess:
				continue
			case compiler.MetaCommandUnrecognizedCommand:
				fmt.Printf("Unrecognized command %q \n", inputBuffer.Buffer)
				continue
			}
		}

		var statement compiler.Statement
		switch compiler.PrepareStatement(inputBuffer, &statement) {
		case compiler.PrepareSuccess:

		case compiler.PrepareUnrecognizedStatement:
			fmt.Printf("Unrecognized command at start of %q \n", inputBuffer.Buffer)
			continue

		case compiler.PrepareSyntaxError:
			fmt.Println("Syntax error. Could not parse statement.")
			continue
		}

		compiler.ExecuteStatement(statement)
		fmt.Println("Executed")

	}
}

func PrintPrompt() {
	fmt.Printf("db > ")
}
