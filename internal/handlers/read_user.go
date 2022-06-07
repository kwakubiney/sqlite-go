package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sqlite-go/internal/engine"
)

func (h *Handler) ReadUser(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		arrayOfRows, err := engine.DeserializeAllRows(h.DB)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to read all user records",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "records have successfully been read",
			"data":    arrayOfRows,
		})
		
	}



	row, err := engine.DeserializeSpecificRow(h.DB, id)
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



func (h *Handler) ReadUsers(c *gin.Context) {
		arrayOfRows, err := engine.DeserializeAllRows(h.DB)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to read all user records",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "records have successfully been read",
			"data":    arrayOfRows,
		})
		
}
