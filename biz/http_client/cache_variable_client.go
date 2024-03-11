package http_client

import (
	"fmt"

	"qinglin.org/cache_variable/biz/model/cache_variable/cache_variable_service"
)

var (
	CacheVariableClient cache_variable_service.Client
)

func InitCacheVariableClient(port int32) error {
	var err error
	if CacheVariableClient, err = cache_variable_service.NewCacheVariableServiceClient(
		fmt.Sprintf("http://127.0.0.1:%d", port),
		cache_variable_service.WithHertzClientOption(),
		cache_variable_service.WithHertzClientMiddleware(),
	); err != nil {
		return err
	}

	return nil
}
