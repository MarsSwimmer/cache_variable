package server_handlers

import (
	"context"
	"fmt"

	"qinglin.org/cache_variable/biz/component/local_cache"
	"qinglin.org/cache_variable/biz/model/cache_variable_server"
)

func GetValue(ctx context.Context, req *cache_variable_server.GetValueReq, resp *cache_variable_server.GetValueResp) error {
	if req.Key == "" {
		return fmt.Errorf("key is empty")
	}

	v, err := local_cache.GetFreeCacheService().Get([]byte(req.Key))
	if err != nil {
		return fmt.Errorf("GetValue error, errMsg:%s", err.Error())
	}

	resp.Value = string(v)
	return nil
}
