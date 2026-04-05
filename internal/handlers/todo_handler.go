package handlers

import (
	"net/http"
	"strconv"

	"github.com/danchenkov/todo-gin/internal/models"
	"github.com/danchenkov/todo-gin/internal/storage"
	"github.com/gin-gonic/gin"
)

func ShowTodos(c *gin.Context) {
	todos, err := storage.LoadTodos()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to load todos")
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"todos": todos,
	})
}

func NewTodoForm(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", nil)
}

func CreateTodo(c *gin.Context) {
	text := c.PostForm("text")

	todos, _ := storage.LoadTodos()

	todo := models.Todo{
		ID:        len(todos) + 1,
		Text:      text,
		Completed: false,
	}

	todos = append(todos, todo)
	storage.SaveTodos(todos)

	c.Redirect(http.StatusFound, "/")
}

func ToggleTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	todos, _ := storage.LoadTodos()

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = !todos[i].Completed
			break
		}
	}

	storage.SaveTodos(todos)
	c.Redirect(http.StatusFound, "/")
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	todos, _ := storage.LoadTodos()
	filtered := []models.Todo{}

	for _, todo := range todos {
		if todo.ID != id {
			filtered = append(filtered, todo)
		}
	}

	storage.SaveTodos(filtered)
	c.Redirect(http.StatusFound, "/")
}
