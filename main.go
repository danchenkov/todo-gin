package todogin
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/yourname/todo-gin/internal/handlers"
)

func main() {
    router := gin.Default()

    router.LoadHTMLGlob("internal/templates/*")
    router.Static("/static", "./static")

    router.GET("/", handlers.ShowTodos)
    router.GET("/todos/new", handlers.NewTodoForm)
    router.POST("/todos", handlers.CreateTodo)
    router.POST("/todos/:id/toggle", handlers.ToggleTodo)
    router.POST("/todos/:id/delete", handlers.DeleteTodo)

    router.Run(":8080")
}