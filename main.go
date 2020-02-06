package main

import (
	"github.com/danrusu/go-chi-todo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var (
	id    int
	todos []*models.Todo
)

func main() {

	todos = []*models.Todo{}

	router := gin.Default()

	v1 := router.Group("/api/todo")
	{
		v1.GET("/", getAllTodos)
		v1.POST("/reset", resetAllTodos)
		v1.GET("/:id", getTodo)
		v1.POST("/", createTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}

	router.Run("localhost:3000")
}

func findTodoInList(context *gin.Context, todoId int) int {
	for index, todo := range todos {
		if todo.Id == todoId {
			return index
		}
	}
	context.JSON(
		http.StatusNotFound,
		gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo was not found",
		})

	return -1
}

func getTodoId(context *gin.Context) int {
	todoId, _ := strconv.Atoi(context.Param("id"))
	return todoId
}

func getAllTodos(context *gin.Context) {
	context.JSON(
		http.StatusOK, // header
		gin.H{ // body
			"status": http.StatusOK,
			"data":   todos,
		})
}

func resetAllTodos(context *gin.Context) {

	todos = []*models.Todo{}

	context.JSON(
		http.StatusOK, // header
		gin.H{
			"status": http.StatusOK,
			"data":   todos,
		})
}

func getTodo(context *gin.Context) {

	todoId := getTodoId(context)

	foundIndex := findTodoInList(context, todoId)
	if foundIndex == -1 {
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   todos[foundIndex],
		})
}

func createTodo(context *gin.Context) {

	var todo models.Todo
	err := context.BindJSON(&todo) //

	todo.Id = id
	id++

	if err != nil {
		log.Fatalln(err)
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid json body",
			})
		return
	}

	todos = append(todos, &todo)

	context.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Todo successfully created",
			"data":    todo,
		})
}

func updateTodo(context *gin.Context) {

	todoId := getTodoId(context)

	var currentTodo models.Todo
	err := context.BindJSON(&currentTodo)

	if err != nil {
		log.Fatalln(err)
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid json body",
			})
		return
	}

	foundIndex := findTodoInList(context, todoId)

	if foundIndex == -1 {
		return
	}

	currentTodo.Id = todoId
	todos[foundIndex] = &currentTodo

	context.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Todo successfully updated",
		})

}

func deleteTodo(context *gin.Context) {

	todoId := getTodoId(context)

	foundIndex := findTodoInList(context, todoId)
	if foundIndex == -1 {
		return
	}

	todos = append(todos[:foundIndex], todos[foundIndex+1:]...)

	context.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Todo successfully deleted",
		})
}
