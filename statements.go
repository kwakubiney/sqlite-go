package sqlitego

import (
	"bytes"
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
	StatementSelectItem
)

type Statement struct {
	RowToInsert Row
	Type        StatementType
	RowToSelect Row
}

var (
	RowsTable       = make([]Row, 0)
	RowsTableBuffer bytes.Buffer
)

func DoMetaCommand(buffer InputBuffer, db *DB) MetaCommandResult {
	if buffer.Buffer == ".exit" {
		db.Close()
		os.Exit(0)
	} else {
		return MetaCommandUnrecognizedCommand
	}
	return MetaCommandSuccess
}

func PrepareStatement(buffer InputBuffer, statement *Statement) PrepareResult {
	bufferArguments := strings.Fields(buffer.Buffer)
	if len(buffer.Buffer) < 6 {
		return PrepareSyntaxError
	}
	if len(buffer.Buffer) > 6 {
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
					statement.RowToInsert.ID = fmt.Sprint(i)

				}
				statement.RowToInsert.Username = bufferArguments[2]
				statement.RowToInsert.Email = bufferArguments[3]
			}
			return PrepareSuccess
		}
	}

	if len(bufferArguments) == 1 && buffer.Buffer == "select" {
		statement.Type = StatementSelect
		return PrepareSuccess
	}

	if bufferArguments[0] == "select" {
		statement.Type = StatementSelectItem
		if len(bufferArguments) > 2 {
			return PrepareSyntaxError
		} else {
			_, err := strconv.Atoi(bufferArguments[1])
			if err != nil {
				log.Printf("%q is not a valid id\n", bufferArguments[1])
				return PrepareSyntaxError
			}
			statement.RowToSelect.ID = bufferArguments[1]
			return PrepareSuccess
		}
	}
	return PrepareUnrecognizedStatement
}

func ExecuteStatement(statement Statement, db *DB) error {
	switch statement.Type {
	case (StatementInsert):
		return SerializeRow(statement.RowToInsert, db)
	case (StatementSelect):
		return DeserializeAllRows(db)
	case (StatementSelectItem):
		_, err := DeserializeSpecificRow(db, statement.RowToSelect.ID)
		return err
	}
	return nil
}
