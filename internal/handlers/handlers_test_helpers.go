package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


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

func DecodeResponse(t *testing.T, response *httptest.ResponseRecorder) map[string]interface{} {
	var responseBody map[string]interface{}
	assert.NoError(t, json.Unmarshal(response.Body.Bytes(), &responseBody))
	return responseBody
}
