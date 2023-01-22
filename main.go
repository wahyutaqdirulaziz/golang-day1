package main

import (
	"net/http"

	"go-basic/controllers"
	"go-basic/middlewares"
	"go-basic/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()
	r := gin.Default()

	public := r.Group("/api")
	public.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	public.POST("/auth/register", controllers.Register)
	public.POST("/auth/login", controllers.Login)

	protected := r.Group("/api/user")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.POST("/todo", controllers.TodoCreate)
	protected.GET("/todo", controllers.TodoList)
	protected.GET("/todobyid", controllers.GetTodobyId)
	protected.POST("/todobyid", controllers.UpdateTodobyId)

	r.Run()
}
