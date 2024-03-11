package client_handlers

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/common/config"
	http_clients "qinglin.org/cache_variable/biz/http_client"
	"qinglin.org/cache_variable/biz/model/cache_variable"
)

func CacheVariableDel(ctx context.Context, key string) error {
	_, _, err := http_clients.CacheVariableClient.Del(ctx, &cache_variable.DelReq{
		Key: key,
	}, config.WithReadTimeout(5*time.Second))
	if err != nil {
		return err
	}
	return nil
}
