package client_handlers

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/common/config"
	http_clients "qinglin.org/cache_variable/biz/http_client"
	"qinglin.org/cache_variable/biz/model/cache_variable"
)

func CacheVariableGet(ctx context.Context, key string) (string, error) {
	resp, _, err := http_clients.CacheVariableClient.Get(ctx, &cache_variable.GetReq{
		Key: key,
	}, config.WithReadTimeout(5*time.Second))
	if err != nil {
		return "", err
	}
	return resp.Value, nil
}
