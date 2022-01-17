package login

import (
	"blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

// @summary 用户注册接口
// @Description 这是一个用户登录接口
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"msg":"bad request"}"
// @Failure 500 {string} json "{"code":500,"msg":"server error"}"
// @Router /signin [post]
func signInHandler(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "bad request"})
		return
	}
	_, err := models.AddUser(json.Username, json.Password, json.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"code": 200, "msg": "ok"})
}

// @Summary 用户登录接口
// @Produce  json
// @Success 200 {string} json "{"code":200 ,"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"msg":"bad request"}"
// @Failure 500 {string} json "{"code":500,"msg":"server error"}"
// @Router /login [post]
func loginHandler(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "bad request"})
		return
	}
	_, err := models.FindUser(json.Username, json.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"code": 200, "msg": "ok"})

}
