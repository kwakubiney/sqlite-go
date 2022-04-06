package compiler

import (
	"fmt"
	"os"
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
)

const (
	StatementInsert StatementType = iota
	StatementSelect
)

type Statement struct {
	Type StatementType
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
		if buffer.Buffer[:6] == "insert" {
		statement.Type = StatementInsert
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
