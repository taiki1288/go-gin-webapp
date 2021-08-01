package main

import (
	"strconv"
	
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
	todo.Status = status 
	db.Save(&todo)
	db.Close()
}

//DB削除
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けません(dbDelete)")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

	dbInit()

	//index
    router.GET("/", func(ctx *gin.Context){
		todos := dbGetAll()
        ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
    })

	//Create
	router.POST("/new", func(ctx *ginContext){
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbInsert(text, status)
		ctx.Redirect(302, "/")
	})

	//Detail
	router.GET("/detail/:id", func(ctx *ginContext){
		n := ctx.Params("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	//Update
	router.POST("/update/:id", func(ctx *gin.Context){
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbUpdate(id, text, status)
		ctx.Redirect(302, "/")
	})

	//削除確認
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	//Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		dbDelete(id)
		ctx.Redirect(302, "/")
	})

    router.Run()
}