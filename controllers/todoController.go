package controllers

import (
	"go-basic/models"
	"go-basic/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoInput struct {
	Title     string `json:"title" binding:"required"`
	Deskripsi string `json:"deskripsi" binding:"required"`
}

func TodoCreate(c *gin.Context) {
	var input TodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Todo{}

	u.Title = input.Title
	u.Deskripsi = input.Deskripsi

	_, err := u.SaveTodo()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo is Created success"})
}

func TodoList(c *gin.Context) {
	_, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetTodo()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": u})

}

func GetTodobyId(c *gin.Context) {
	_, err := token.ExtractTokenID(c)
	id := c.Query("id")
	intVar, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetTodobyId(uint(intVar))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": u})

}


type TodoUpdate struct {
	Title     string `json:"title" binding:"required"`
	Deskripsi string `json:"deskripsi" binding:"required"`
}
func UpdateTodobyId(c *gin.Context) {
	var input TodoUpdate
	id := c.Query("id")
	intVarb, err := strconv.Atoi(id)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Todo{}

	u.Title = input.Title
	u.Deskripsi = input.Deskripsi

	 u.UpdateTodobyId(uint(intVarb))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo is Created success"})

}
