package local_cache

import "qinglin.org/cache_variable/biz/component/local_cache/impl"

var (
	// 保证单列
	free_cache_service LocalCacheService
)

func Init(cacheSize int32) {
	initFreeCacheService(cacheSize)
}

func initFreeCacheService(cacheSize int32) {
	free_cache_service = impl.NewFreeCache(int(cacheSize))
}

func GetFreeCacheService() LocalCacheService {
	return free_cache_service
}
