package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sqlite-go"
)

func (h *Handler) ReadUser(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse request",
		})
		return
	}

	row, err := sqlitego.DeserializeSpecificRow(h.DB, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to read user's record",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "record has successfully been read",
		"data":    row,
	})
}
