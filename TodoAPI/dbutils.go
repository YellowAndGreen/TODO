package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Sort    int    `json:"sort"`
	Done	bool	`json:"done"`
}

func initdb() {
	//开启数据库
	db, _ := gorm.Open(sqlite.Open("./todo.db"), &gorm.Config{})
	// 自动迁移
	err := db.AutoMigrate(&Todo{})
	if err != nil {
		panic(err)
	}
}

func addTodo(db *gorm.DB, t *Todo) {
	db.Create(&t)
}

func deleteTodo(db *gorm.DB, id int) {
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
}

// 传入已经找到且更新参数的Todo
func updateTodo(db *gorm.DB, t *Todo, id int) {
	var todo Todo
	db.First(&todo, id)
	db.Model(&todo).Updates(map[string]interface{}{"title": t.Title, "content": t.Content, "sort": t.Sort})
}

//func main() {
//	//开启数据库
//	db, _ := gorm.Open(sqlite.Open("./todo.db"), &gorm.Config{})
//	// 添加操作
//	//t := &Todo{Title: "foo2", Content: "bar2", Sort: 1}
//	//addTodo(db, t)
//
//	// 删除
//	//deleteTodo(db, 2)
//
//	// 更新todo
//	t := &Todo{Title: "foo3", Content: "bar3", Sort: 1}
//	updateTodo(db, t, 1)
//}
