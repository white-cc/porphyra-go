package article

import "github.com/gin-gonic/gin"

func LoadArticle(e *gin.Engine) {
	e.GET("/article", getAllArticleHandler)
	e.GET("/article/{id}", getArticleHandler)
	e.POST("/article", postArticleHandler)
	e.DELETE("/article", deleteArticleHandler)
}
