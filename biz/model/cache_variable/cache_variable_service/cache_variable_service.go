// Code generated by hertz generator.

package cache_variable_service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	cache_variable "qinglin.org/cache_variable/biz/model/cache_variable"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type Client interface {
	Get(context context.Context, req *cache_variable.GetReq, reqOpt ...config.RequestOption) (resp *cache_variable.GetResp, rawResponse *protocol.Response, err error)

	Set(context context.Context, req *cache_variable.SetReq, reqOpt ...config.RequestOption) (resp *cache_variable.SetResp, rawResponse *protocol.Response, err error)

	Del(context context.Context, req *cache_variable.DelReq, reqOpt ...config.RequestOption) (resp *cache_variable.DelResp, rawResponse *protocol.Response, err error)
}

type CacheVariableServiceClient struct {
	client *cli
}

func NewCacheVariableServiceClient(hostUrl string, ops ...Option) (Client, error) {
	opts := getOptions(append(ops, withHostUrl(hostUrl))...)
	cli, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	return &CacheVariableServiceClient{
		client: cli,
	}, nil
}

func (s *CacheVariableServiceClient) Get(context context.Context, req *cache_variable.GetReq, reqOpt ...config.RequestOption) (resp *cache_variable.GetResp, rawResponse *protocol.Response, err error) {
	httpResp := &cache_variable.GetResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"key": req.GetKey(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/get_value")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *CacheVariableServiceClient) Set(context context.Context, req *cache_variable.SetReq, reqOpt ...config.RequestOption) (resp *cache_variable.SetResp, rawResponse *protocol.Response, err error) {
	httpResp := &cache_variable.SetResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"key":    req.GetKey(),
			"value":  req.GetValue(),
			"expire": req.GetExpire(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/set_value")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *CacheVariableServiceClient) Del(context context.Context, req *cache_variable.DelReq, reqOpt ...config.RequestOption) (resp *cache_variable.DelResp, rawResponse *protocol.Response, err error) {
	httpResp := &cache_variable.DelResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"key": req.GetKey(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/delete_value")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

var defaultClient, _ = NewCacheVariableServiceClient("http://127.0.0.1:6366")

func ConfigDefaultClient(ops ...Option) (err error) {
	defaultClient, err = NewCacheVariableServiceClient("http://127.0.0.1:6366", ops...)
	return
}

func Get(context context.Context, req *cache_variable.GetReq, reqOpt ...config.RequestOption) (resp *cache_variable.GetResp, rawResponse *protocol.Response, err error) {
	return defaultClient.Get(context, req, reqOpt...)
}

func Set(context context.Context, req *cache_variable.SetReq, reqOpt ...config.RequestOption) (resp *cache_variable.SetResp, rawResponse *protocol.Response, err error) {
	return defaultClient.Set(context, req, reqOpt...)
}

func Del(context context.Context, req *cache_variable.DelReq, reqOpt ...config.RequestOption) (resp *cache_variable.DelResp, rawResponse *protocol.Response, err error) {
	return defaultClient.Del(context, req, reqOpt...)
}
