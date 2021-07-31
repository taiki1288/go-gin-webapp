package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"

)

type Todo struct {
	gorm.Model
	Text string
	Status string
}

//DB初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けません(dbInit)")
	}
	db.AutoMigrate(Todo{})
	defer db.Close()
}

func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けません(dbInsert)")
	}

	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

//DB取得
func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けません(dbGetAll())")
	}

	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けません(dbGetOne())")
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

//DB更新
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けません(dbUpdate)")
	}
	var todo Todo 
	db.First(&todo, id)
	todo.Text = text 
	todo.status = status 
	db.Save(&todo)
	db.Close()
}

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

	data := "Hello Go/Gin!!"

    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(200, "index.html", gin.H{"data": data})
    })

    router.Run()
}