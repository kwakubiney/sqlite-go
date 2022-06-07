package engine_test

import (
	"log"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/sqlite-go/internal/engine"
	"github.com/sqlite-go"
)

func TestPrepareStatement(t *testing.T) {
	type PrepareStatementTestSuite struct {
		Buffer    sqlitego.InputBuffer
		Statement engine.Statement
		Response  engine.PrepareResult
	}
	for _, buffer := range []PrepareStatementTestSuite{
		{
			sqlitego.InputBuffer{
				Buffer: "insert 1 adam kwaku@mail.com",
			},

			engine.Statement{},

			engine.PrepareSuccess,
		},
		{
			sqlitego.InputBuffer{
				Buffer: "insert adam kwaku kwaku@mail.com",
			},

			engine.Statement{},
			engine.PrepareSyntaxError,
		},
		{
			sqlitego.InputBuffer{
				Buffer: "some random unrecognizable statement",
			},
			engine.Statement{},
			engine.PrepareUnrecognizedStatement,
		}} {

		actualResponse := engine.PrepareStatement(buffer.Buffer, &buffer.Statement)
		assert.Equal(t, buffer.Response, actualResponse)
	}
	DB.TestClose()
}

func BenchmarkSerialization(b *testing.B) {
	row := engine.Row{
		ID:       "5",
		Username: "kwakubiney",
		Email:    "k@mail.com",
	}
	for n := 0; n < b.N; n++ {
		err := engine.SerializeRow(row, DB)
		if err != nil {
			log.Fatalln(err)
		}
	}

	err := DB.TestClose()
	if err != nil {
		log.Println(err)
	}
}

