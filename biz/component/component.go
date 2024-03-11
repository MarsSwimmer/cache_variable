package component

import "qinglin.org/cache_variable/biz/component/local_cache"

func InitComponent(cacheSize int32){
	local_cache.Init(cacheSize)
}