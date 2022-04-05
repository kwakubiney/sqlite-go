package main

import (
	"bufio"
	"fmt"
	"os"
)

type InputBuffer struct {
	Buffer       string
	BufferLength int
	InputLength  int
}

func NewInputBuffer() InputBuffer {
	newInputBuffer := InputBuffer{}
	newInputBuffer.Buffer = ""
	newInputBuffer.BufferLength = 16
	newInputBuffer.InputLength = 16
	return newInputBuffer
}

func main() {
	inputBuffer := NewInputBuffer()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		command := scanner.Text()
		inputBuffer.Buffer = command
		if len(command) <= 1 {
			fmt.Println("Error reading input")
		}

		if command == "exit()" && len(command) > 1{
			os.Exit(0)
		} else {
			fmt.Printf("Unrecognized command '%s'.\n", inputBuffer.Buffer)
		}
	}

}
