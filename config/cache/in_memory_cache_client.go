package cache

import (
	"time"

	"github.com/Adgytec/adgytec-flow/utils/core"
	lru "github.com/hashicorp/golang-lru/v2/expirable"
)

const (
	defaultCacheSize = 1 << 13 // 8192
	defaultCacheTTL  = 5 * time.Minute
)

type inMemoryLruCache struct {
	cache *lru.LRU[string, []byte]
}

func (cc *inMemoryLruCache) Get(key string) ([]byte, bool) {
	return cc.cache.Get(key)
}

func (cc *inMemoryLruCache) Set(key string, data []byte) {
	cc.cache.Add(key, data)
}

func (cc *inMemoryLruCache) Delete(key string) {
	cc.cache.Remove(key)
}

func CreateInMemoryCacheClient() core.ICacheClient {
	return &inMemoryLruCache{
		cache: lru.NewLRU[string, []byte](defaultCacheSize, nil, defaultCacheTTL),
	}
}
