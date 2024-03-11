package server_handlers

import (
	"context"
	"fmt"

	"qinglin.org/cache_variable/biz/component/local_cache"
	"qinglin.org/cache_variable/biz/model/cache_variable_server"
)

func SetValue(ctx context.Context, req *cache_variable_server.SetValueReq, resp *cache_variable_server.SetValueResp) error {

	if req.Key == "" {
		return fmt.Errorf("key is empty")
	}

	if req.Expire <= 0 {
		return fmt.Errorf("expire:[%d] is invalid", req.Expire)
	}

	return local_cache.GetFreeCacheService().Set([]byte(req.Key), []byte(req.Value), int(req.Expire))
}
