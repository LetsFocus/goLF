package goLF

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/LetsFocus/goLF/errors"
	logger2 "github.com/LetsFocus/goLF/logger"
	"io"
	"net/http"
	"strings"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const interval = 5

type Client struct {
	*http.Client
	url           string
	logger        *logger2.CustomLogger
	headerKeys    []string
	customHeaders map[string]string
}

type HTTPResponse struct {
	Body       []byte
	StatusCode int
	headers    http.Header
}

func NewClient(resourceAddr string, logger *logger2.CustomLogger) *Client {
	if resourceAddr == "" {
		logger.Info("value for resourceAddress is empty")
	} else {
		resourceAddr = strings.TrimRight(resourceAddr, "/")
	}

	transport := otelhttp.NewTransport(http.DefaultTransport)

	httpSvc := &Client{
		url:    resourceAddr,
		logger: logger,
		Client: &http.Client{Transport: transport, Timeout: interval * time.Second},
	}

	return httpSvc
}

func (c *Client) Call(ctx context.Context, method, target string, params map[string]interface{}, body []byte, headers map[string]string) (HTTPResponse, error) {
	target = strings.TrimLeft(target, "/")
	correlationID, _ := ctx.Value("correlationID").(string)
	c.logger.Infof("correlationID for the request is %v", correlationID)

	req, err := c.createRequest(ctx, method, target, params, body, headers)
	if err != nil {
		return HTTPResponse{}, err
	}

	var resp *http.Response
	resp, err = c.Do(req)
	if err != nil {
		return HTTPResponse{}, err
	}

	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		return HTTPResponse{}, err
	}

	return HTTPResponse{StatusCode: resp.StatusCode, Body: byteData, headers: resp.Header}, nil
}

func (c *Client) createRequest(ctx context.Context, method, target string, params map[string]interface{}, body []byte, headers map[string]string) (*http.Request, error) {
	uri := c.url + "/" + target

	if target == "" {
		uri = c.url
	}

	req, err := http.NewRequestWithContext(ctx, method, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Errors{StatusCode: http.StatusInternalServerError}
	}

	req.Header.Add("content-type", "application/json")

	for k, v := range c.customHeaders {
		req.Header.Set(k, v)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if (method == http.MethodGet || method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch) && params != nil {
		setQueryParams(req, params)
	}

	return req, nil
}

func setQueryParams(r *http.Request, params map[string]interface{}) {
	q := r.URL.Query()

	for k, v := range params {
		switch vt := v.(type) {
		case []string:
			for _, val := range vt {
				q.Add(k, val)
			}
		default:
			q.Set(k, fmt.Sprintf("%v", v))
		}
	}

	r.URL.RawQuery = q.Encode()
}

func (c *Client) Bind(data []byte, i interface{}) error {
	err := json.Unmarshal(data, i)
	if err != nil {
		return err
	}

	return nil
}
