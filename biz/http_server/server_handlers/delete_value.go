package server_handlers

import (
	"context"
	"fmt"

	"qinglin.org/cache_variable/biz/component/local_cache"
	"qinglin.org/cache_variable/biz/model/cache_variable_server"
)

func DeleteValue(ctx context.Context, req *cache_variable_server.DeleteValueReq, resp *cache_variable_server.DeleteValueResp) error {
	if affected := local_cache.GetFreeCacheService().Del([]byte(req.Key)); affected {
		return nil
	}

	return fmt.Errorf("key:%s is not exist", req.Key)
}
