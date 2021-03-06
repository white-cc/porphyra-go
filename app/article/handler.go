package article

import (
	"blog/models"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Name   string `json:"name" binding:"required"`
	Author string `json:"author" binding:"required"`
	Text   string `json:"text" binding:"required"`
	Tag    string `json:"tag"`
}

type Article_id struct {
	ID string `uri:"id" binding:"required"`
}

// @summary 获取单个文章
// @Router /article/:id [get]
func getArticleHandler(c *gin.Context) {
	var id Article_id
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": "bad request"})
		return
	}
	article, err := models.GetArticle(id.ID)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "server error"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "data": article})

}

// @summary 获取所有文章
// @Failure 500 {string} json "{"code":500,"msg":"server error"}"
// @Router /article [get]
func getAllArticleHandler(c *gin.Context) {
	data, err := models.GetAllArticle()
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "server error"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "data": data})
}

// @summary 提交文章
// @Router /article [post]
func postArticleHandler(c *gin.Context) {
	var json Article
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": "BadRequest"})
		return
	}
	id, err := models.AddAticle(json.Name, json.Author, json.Text, json.Tag)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "server error"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "id": id})
}

// @summary 删除文章
// @Router /article/:id [delete]
func deleteArticleHandler(c *gin.Context) {
	var id Article_id
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": "bad request"})
		return
	}
	if err := models.DeleteArticle(id.ID); err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "server error"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "ok"})
}
