package sqlitego

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func (db *DB) TestClose() error {
	if !db.Opened {
		return nil
	}
	db.Opened = false

	RemoveIndexFile(db)
	RemoveDBTestFile()

	db.Path = ""
	db.IndexFilePath = ""

	if err := db.File.Close(); err != nil {
		return fmt.Errorf("failed to close db: %s", err)
	}

	if err := db.IndexFile.Close(); err != nil {
		return fmt.Errorf("failed to close index file: %s", err)
	}

	return nil
}

//TODO: This defeats DRY principle, have to create a general function to truncate files based on file names.
func RemoveDBTestFile() {
	if err := os.Truncate("test-db", 0); err != nil {
		log.Printf("failed to truncate: %v", err)
	}
}

func MakeTestRequest(t *testing.T, route string, body interface{}, method string) *http.Request {
	reqBody, err := json.Marshal(body)
	assert.NoError(t, err)

	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest(method, route, bytes.NewReader(reqBody))
	assert.NoError(t, err)

	return req
}

func BootstrapServer(req *http.Request, routeHandlers *gin.Engine) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	routeHandlers.ServeHTTP(responseRecorder, req)
	return responseRecorder
}