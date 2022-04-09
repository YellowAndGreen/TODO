package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// db作为全局变量
var db *gorm.DB

// GetTodo 获取单条todo
func GetTodo(c *gin.Context) {
	var todo Todo
	//获取get请求的参数
	id := c.Param("todo_id")
	result := db.First(&todo, id)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": todo,
		})
	}
}

// CreateTodo 添加一个Todo
func CreateTodo(c *gin.Context) {
	var todo Todo
	// 将请求体的数据解析为station
	if err := c.BindJSON(&todo); err == nil {
		db.Create(&todo)
		c.JSON(http.StatusOK, gin.H{
			"result": "OK",
		})
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

// RemoveTodo 删除一个todo
func RemoveTodo(c *gin.Context) {
	var todo Todo
	//获取get请求的参数
	id := c.Param("todo_id")

	result := db.First(&todo, id)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
	} else {
		db.Delete(&todo)
		c.JSON(200, gin.H{
			"result": todo,
		})
	}
}

// GetAllTodo 获取当天所有的todo
func GetAllTodo(c *gin.Context) {
	var todos []Todo
	now := time.Now()
	result := db.Where("updated_at > ?", now.Format(time.RFC3339)).Find(&todos)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": todos,
		})
	}
}

// UpdateTodo 更新Todo参数
func UpdateTodo(c *gin.Context)  {

}

func main() {
	initdb()
	// db赋值
	db, _ = gorm.Open(sqlite.Open("./todo.db"), &gorm.Config{})
	r := gin.Default()
	// Add routes to REST verbs
	r.GET("/api/todos/:todo_id", GetTodo)
	r.GET("/api/todos/today", GetAllTodo)
	r.GET("/api/todos/remove/:todo_id", RemoveTodo)
	r.POST("/api/todos/create", CreateTodo)
	r.Run(":9000")
}
