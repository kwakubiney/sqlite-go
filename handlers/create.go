package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sqlite-go"
)

type CreateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (h *Handler) CreateUser(c *gin.Context) {
	var cu CreateUserRequest

	err := c.BindJSON(&cu)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse request",
		})
	}

	row := &sqlitego.Row{
		ID:       cu.ID,
		Username: cu.Username,
		Email:    cu.Email,
	}

	err = sqlitego.SerializeRow(*row, h.DB)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create record",
		})
	}

	log.Println(row.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "record has successfully been created",
		"data":    row,
	})
}
