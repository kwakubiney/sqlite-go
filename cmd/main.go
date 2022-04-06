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
		scanner.Scan()
		command := scanner.Text()
		inputBuffer.Buffer = command

		if strings.HasPrefix(inputBuffer.Buffer, "."){
			switch compiler.DoMetaCommand(inputBuffer){
			case compiler.MetaCommandSuccess:
				continue;
			case compiler.MetaCommandUnrecognizedCommand:
				fmt.Printf("Unrecognized command %s \n", inputBuffer.Buffer)
				continue;
			}}


		var statement compiler.Statement
			switch compiler.PrepareStatement(inputBuffer, &statement){
			case compiler.PrepareSuccess:
				
			case compiler.PrepareUnrecognizedStatement:
				fmt.Printf("Unrecognized command at start of %s \n", inputBuffer.Buffer)
				continue;
		}

		compiler.ExecuteStatement(statement)
		fmt.Println("Executed")
			
		}
		
	}
