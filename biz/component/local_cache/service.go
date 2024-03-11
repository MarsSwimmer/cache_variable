package local_cache

type LocalCacheService interface {
	// 设置k-v
	//
	// @param
	// key: 缓存的key
	// value: 缓存的value
	// expire: 过期时间，内存不够时会被覆盖写，单位秒
	//
	// @return
	// err: 设置异常
	Set(key, value []byte, expire int) error

	// 获取key
	//
	// @param
	// key: 缓存key
	//
	// @return
	// value: 获取缓存的值
	// err: 读取缓存异常
	Get(key []byte) ([]byte, error)

	// 删除key-value
	//
	// @param
	// key: 缓存key
	//
	// @return
	// bool: 是否删除成功
	Del(key []byte) bool
}
