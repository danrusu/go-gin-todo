package main

import (
	"github.com/danrusu/go-chi-todo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	id    int
	todos []*models.Todo
)

func main() {
	port := os.Getenv("PORT")

	todos = []*models.Todo{}

	router := gin.Default()

	router.LoadHTMLGlob("./html/*.html")
	router.StaticFile("favicon.ico", "./html/favicon.ico")
	router.StaticFile("todo.css", "./html/todo.css")
	router.StaticFile("todo.js", "./html/todo.js")
	//	router.Static("/static", "static")

	todoApiRouter := router.Group("/api/todo")
	{
		todoApiRouter.GET("/healthcheck", healthcheck)
		todoApiRouter.GET("/", getAllTodos)
		todoApiRouter.POST("/reset", resetAllTodos)
		todoApiRouter.GET("/:id", getTodo)
		todoApiRouter.POST("/", createTodo)
		todoApiRouter.PUT("/:id", updateTodo)
		todoApiRouter.DELETE("/:id", deleteTodo)
	}

	uiRouter := router.Group("/")
	{
		uiRouter.GET("/", func(context *gin.Context) {
			context.HTML(http.StatusOK, "index.html", nil)
		})
	}

	router.Run(":" + port)
}

func healthcheck(context *gin.Context){
	context.JSON(
        	http.StatusOK,
                gin.H{
                        "status": "healthy",
                })
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

func getAllTodos(context *gin.Context) {
	context.JSON(
		http.StatusOK, // header
		gin.H{ // body
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

func resetAllTodos(context *gin.Context) {

	todos = []*models.Todo{}
	id = 0

	context.JSON(
		http.StatusOK, // header
		gin.H{
			"status": http.StatusOK,
			"data":   todos,
		})
}

func getTodoId(context *gin.Context) int {
	todoId, _ := strconv.Atoi(context.Param("id"))
	return todoId
}
