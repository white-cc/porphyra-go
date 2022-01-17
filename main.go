package main

import (
	"blog/app/article"
	"blog/app/login"
	_ "blog/docs"
	"blog/models"
	"blog/routers"
	"log"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title blog后端接口
// @version 0.1
// @description 个人博客接口列表
func main() {
	godotenv.Load()
	routers.Include(login.Loadlogin)
	routers.Include(article.LoadArticle)
	//TODO
	// routers.Include(comment.LoadComment)
	models.Init()

	r := routers.Init()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
