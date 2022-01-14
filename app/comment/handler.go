package comment

import "github.com/gin-gonic/gin"

type Comment struct {
	Article_id string
	Comment_id string
	Text       string
}

// @summary 获取文章评论
// @Router /comment [get]
func getCommentHandler(c *gin.Context) {

}

// @summary 提交文章评论
// @Router /comment [post]
func postCommentHandler(c *gin.Context) {

}

// @summary 删除文章评论
// @Router /comment [delete]
func deleteCommentHandler(c *gin.Context) {

}
