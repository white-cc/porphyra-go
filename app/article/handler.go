package article

import "github.com/gin-gonic/gin"

type Article struct {
	Text string
}

// @summary 获取单个文章
// @Router /article/{id} [get]
func getArticleHandler(c *gin.Context) {

}

// @summary 获取所有文章大纲
// @Router /article [get]
func getAllArticleHandler(c *gin.Context) {

}

// @summary 提交文章
// @Router /article [post]
func postArticleHandler(c *gin.Context) {

}

// @summary 删除文章
// @Router /article [delete]
func deleteArticleHandler(c *gin.Context) {

}
