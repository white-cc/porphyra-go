package comment

import "github.com/gin-gonic/gin"

func LoadComment(e *gin.Engine) {
	e.GET("/comment", getCommentHandler)
	e.POST("/comment", postCommentHandler)
	e.DELETE("/comment", deleteCommentHandler)
}
