package service

import "github.com/gin-gonic/gin"

type service interface {
	Get(ctx *gin.Context, target string, params map[string]interface{}, headers map[string]string) (HTTPResponse, error)
	Post(ctx *gin.Context, target string, body []byte, headers map[string]string) (HTTPResponse, error)
	Put(ctx *gin.Context, target string, body []byte, params map[string]interface{}, headers map[string]string) (HTTPResponse, error)
	Patch(ctx *gin.Context, target string, body []byte, params map[string]interface{}, headers map[string]string) (HTTPResponse, error)
	Delete(ctx *gin.Context, target string, params map[string]interface{}, headers map[string]string) (HTTPResponse, error)
	Bind(data []byte, i interface{}) error
}
