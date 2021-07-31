package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

)

type Todo struct {
	gorm.Model
	Text string
	Status string
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