package routers

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

func Include(opt ...Option) {
	options = append(options, opt...)
}

func Init() *gin.Engine {
	e := gin.New()
	for _, opt := range options {
		opt(e)
	}
	return e
}
