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

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

	data := "Hello Go/Gin!!"

    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(200, "index.html", gin.H{"data": data})
    })

    router.Run()
}