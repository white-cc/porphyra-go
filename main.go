package main

import (
	"blog/app/article"
	"blog/app/comment"
	"blog/app/login"
	_ "blog/docs"
	"blog/routers"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title blog后端接口
// @version 0.1
// @description 个人博客接口列表
func main() {
	routers.Include(login.Loadlogin)
	routers.Include(article.LoadArticle)
	routers.Include(comment.LoadComment)

	r := routers.Init()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
