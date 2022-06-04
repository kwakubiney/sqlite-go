package sqlitego_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"github.com/gin-gonic/gin"

	"github.com/sqlite-go"
	"github.com/stretchr/testify/assert"
	"github.com/sqlite-go/handlers"
	"github.com/sqlite-go/server"
)

var db *sqlitego.DB
var routeHandlers *gin.Engine

func TestMain(m *testing.M) {
	var err error
	db, err = sqlitego.DbOpen("test-db", "test-index", 0644)
	if err != nil {
		log.Println(err)
	}
	handlers := handlers.New(db)
	server := server.New(handlers)
	routeHandlers = server.SetupRoutes()
	os.Exit(m.Run())
}

func TestSerializationAndDeserialization(t *testing.T) {
	type DeserializeCases struct {
		Row              sqlitego.Row
		ExpectedResponse string
	}

	for _, row := range []DeserializeCases{
		{
			sqlitego.Row{ID: "5",
				Username: "Kwaku Egyir Biney",
				Email:    "kwakz@yahoo.com",
			},
			"ID: 5, Name: Kwaku Egyir Biney, Email: kwakz@yahoo.com"},
		{
			sqlitego.Row{ID: "10",
				Username: "Sally Akua",
				Email:    "sally@yahoo.com",
			},
			"ID: 10, Name: Sally Akua, Email: sally@yahoo.com"},
	} {
		err := sqlitego.SerializeRow(row.Row, db)
		assert.NoError(t, err)

		deserializedRow, err := sqlitego.DeserializeSpecificRow(db, row.Row.ID)

		assert.NoError(t, err)

		assert.Equal(t, fmt.Sprintf("ID: %s, Name: %s, Email: %s", row.Row.ID, row.Row.Username, row.Row.Email), deserializedRow)
	}

	db.TestClose()
}


