package engine_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sqlite-go/internal/engine"
	"github.com/sqlite-go/internal/handlers"
	"github.com/sqlite-go/internal/server"
	"github.com/stretchr/testify/assert"
)

var DB *engine.DB
var routeHandlers *gin.Engine

func TestSerializationAndDeserialization(t *testing.T) {
	DB, err := engine.DbOpen("test-db", "test-index", 0644)
	if err != nil {
		log.Println(err)
	}
	handlers := handlers.New(DB)
	server := server.New(handlers)
	routeHandlers = server.SetupRoutes()
	
	type DeserializeCases struct {
		Row              engine.Row
		ExpectedResponse string
	}

	for _, row := range []DeserializeCases{
		{
			engine.Row{ID: "5",
				Username: "Kwaku Egyir Biney",
				Email:    "kwakz@yahoo.com",
			},
			"ID: 5, Name: Kwaku Egyir Biney, Email: kwakz@yahoo.com"},
		{
			engine.Row{ID: "10",
				Username: "Sally Akua",
				Email:    "sally@yahoo.com",
			},
			"ID: 10, Name: Sally Akua, Email: sally@yahoo.com"},
	} {
		err := engine.SerializeRow(row.Row, DB)
		assert.NoError(t, err)

		deserializedRow, err := engine.DeserializeSpecificRow(DB, row.Row.ID)

		assert.NoError(t, err)

		assert.Equal(t, fmt.Sprintf("ID: %s, Name: %s, Email: %s", row.Row.ID, row.Row.Username, row.Row.Email), deserializedRow)
	}

	handlers.DB.TestClose()
}


