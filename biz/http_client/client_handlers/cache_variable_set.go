package client_handlers

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/common/config"
	http_clients "qinglin.org/cache_variable/biz/http_client"
	"qinglin.org/cache_variable/biz/model/cache_variable"
)

func CacheVariableSet(ctx context.Context, key, value string, expire int32) error {
	_, _, err := http_clients.CacheVariableClient.Set(ctx, &cache_variable.SetReq{
		Key:    key,
		Value:  value,
		Expire: expire,
	}, config.WithReadTimeout(5*time.Second))
	if err != nil {
		return err
	}

	return nil
}
