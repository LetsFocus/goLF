package service

import (
	"context"
)

type service interface {
	Get(ctx context.Context, target string, params map[string]interface{}, headers map[string]string) (HTTPResponse, error)
	Post(ctx context.Context, target string, body []byte, headers map[string]string) (HTTPResponse, error)
	Put(ctx context.Context, target string, body []byte, params map[string]interface{}, headers map[string]string) (HTTPResponse, error)
	Patch(ctx context.Context, target string, body []byte, params map[string]interface{}, headers map[string]string) (HTTPResponse, error)
	Delete(ctx context.Context, target string, params map[string]interface{}, headers map[string]string) (HTTPResponse, error)
	Bind(data []byte, i interface{}) error
}
