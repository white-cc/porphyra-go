package login

import "github.com/gin-gonic/gin"

type Login struct {
	User     string
	Password string
	Email    string
}

// @summary 用户注册接口
// @Description 这是一个用户登录接口
// @Accept json
// @Produce json
// @Param name query string true "Name"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 404 {string} json "{"code":404,"data":{},"msg":"404 not find"}"
// @Failure 500 {string} json "{"code":500,"data":{},"msg":"server error"}"
// @Router /signin [post]
func signInHandler(c *gin.Context) {

}

// @Summary 用户登录接口
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /login [post]
func loginHandler(c *gin.Context) {

}
