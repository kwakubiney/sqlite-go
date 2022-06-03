package sqlitego_test

import (
	"log"
	"testing"

	"github.com/sqlite-go"
	"github.com/stretchr/testify/assert"
)

func TestPrepareStatement(t *testing.T) {
	type PrepareStatementTestSuite struct {
		Buffer    sqlitego.InputBuffer
		Statement sqlitego.Statement
		Response  sqlitego.PrepareResult
	}
	for _, buffer := range []PrepareStatementTestSuite{
		{
			sqlitego.InputBuffer{
				Buffer: "insert 1 adam kwaku@mail.com",
			},

			sqlitego.Statement{},

			sqlitego.PrepareSuccess,
		},
		{
			sqlitego.InputBuffer{
				Buffer: "insert adam kwaku kwaku@mail.com",
			},

			sqlitego.Statement{},
			sqlitego.PrepareSyntaxError,
		},
		{
			sqlitego.InputBuffer{
				Buffer: "some random unrecognizable statement",
			},
			sqlitego.Statement{},
			sqlitego.PrepareUnrecognizedStatement,
		}} {

		actualResponse := sqlitego.PrepareStatement(buffer.Buffer, &buffer.Statement)
		assert.Equal(t, buffer.Response, actualResponse)
	}
	db.TestClose()
}

func BenchmarkSerialization(b *testing.B) {
	row := sqlitego.Row{
		ID:       "5",
		Username: "kwakubiney",
		Email:    "k@mail.com",
	}
	for n := 0; n < b.N; n++ {
		err := sqlitego.SerializeRow(row, db)
		if err != nil {
			log.Fatalln(err)
		}
	}

	err := db.TestClose()
	if err != nil {
		log.Println(err)
	}
}
