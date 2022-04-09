package compiler

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MetaCommandResult int
type PrepareResult int
type StatementType int

const (
	MetaCommandSuccess MetaCommandResult = iota
	MetaCommandUnrecognizedCommand
)

const (
	PrepareSuccess PrepareResult = iota
	PrepareUnrecognizedStatement
	PrepareSyntaxError
)

const (
	StatementInsert StatementType = iota
	StatementSelect
)

type Statement struct {
	RowToInsert Row
	Type        StatementType
}

func DoMetaCommand(buffer InputBuffer) MetaCommandResult {
	if buffer.Buffer == ".exit" {
		os.Exit(0)
	} else {
		return MetaCommandUnrecognizedCommand
	}
	return MetaCommandSuccess
}

func PrepareStatement(buffer InputBuffer, statement *Statement) PrepareResult {
	if len(buffer.Buffer) > 6 {
		bufferArguments := strings.Fields(buffer.Buffer)
		if bufferArguments[0] == "insert" {
			statement.Type = StatementInsert
			if len(bufferArguments) < 4 {
				return PrepareSyntaxError
			} else {
				i, err := strconv.Atoi(bufferArguments[1])
				if err != nil {
					log.Printf("%q is not a valid id\n", bufferArguments[1])
					return PrepareSyntaxError
				} else {
					statement.RowToInsert.id = int32(i)
				}
				statement.RowToInsert.username = bufferArguments[2]
				statement.RowToInsert.username = bufferArguments[3]
			}
			return PrepareSuccess
		}
	}

	if buffer.Buffer == "select" {
		statement.Type = StatementSelect
		return PrepareSuccess
	}

	return PrepareUnrecognizedStatement
}

func ExecuteStatement(statement Statement) {
	switch statement.Type {
	case (StatementInsert):
		fmt.Println("This is where an insert will be done.")
	case (StatementSelect):
		fmt.Println("This is where a select will be done.")

	}
}
