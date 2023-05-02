package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "1", Item: "Read Book", Completed: false},
	{ID: "1", Item: "Record Video", Completed: false},
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}

	}
	return nil, errors.New("todos not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})

	}
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodos(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, todos)

}

func main() {

	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:7676")

}
