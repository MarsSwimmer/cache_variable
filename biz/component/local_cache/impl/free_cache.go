package impl

import "github.com/coocood/freecache"

type FreeCache struct{
	cache freecache.Cache
}

func NewFreeCache(cacheSize int) *FreeCache{
	return &FreeCache{
		cache: *freecache.NewCache(cacheSize),
	}
}

func (s *FreeCache) Set(key, value []byte, expire int) error{
	return s.cache.Set(key, value, expire)
}

func (s *FreeCache) Get(key []byte) ([]byte, error){
	return s.cache.Get(key)
}

func (s *FreeCache) Del(key []byte) bool{
	return s.cache.Del(key)
}