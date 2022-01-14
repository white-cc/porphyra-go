package login

import "github.com/gin-gonic/gin"

func Loadlogin(e *gin.Engine) {
	e.POST("/login", loginHandler)
	e.POST("/signin", signInHandler)
}
