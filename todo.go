package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json: "id"`
	Item      string `json: "item"`
	Completed bool   `json: "completed"`
}

var todos = []todo{
	{ID: "1", Item: "Pay Bills", Completed: false},
	{ID: "2", Item: "Feed Dog", Completed: false},
	{ID: "3", Item: "Buy Groceries", Completed: false},
	{ID: "4", Item: "Exercise", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:7766")

}
