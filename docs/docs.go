// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/article": {
            "get": {
                "summary": "获取所有文章",
                "responses": {}
            },
            "post": {
                "summary": "提交文章",
                "responses": {}
            },
            "delete": {
                "summary": "删除文章",
                "responses": {}
            }
        },
        "/comment": {
            "get": {
                "summary": "获取文章评论",
                "responses": {}
            },
            "post": {
                "summary": "提交文章评论",
                "responses": {}
            },
            "delete": {
                "summary": "删除文章评论",
                "responses": {}
            }
        },
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录接口",
                "responses": {
                    "200": {
                        "description": "{\"code\":200 ,\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"code\":400,\"msg\":\"bad request\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"code\":500,\"msg\":\"server error\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "description": "这是一个用户登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户注册接口",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"code\":400,\"msg\":\"bad request\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"code\":500,\"msg\":\"server error\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "blog后端接口",
	Description: "个人博客接口列表",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
