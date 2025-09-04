package cache

import (
	"fmt"
	"log"

	"github.com/Adgytec/adgytec-flow/config/serializer"
	"golang.org/x/sync/singleflight"
)

type Cache[T any] interface {
	Get(string, func() (T, error)) (T, error)
	Delete(string)
}

type CacheClient interface {
	Get(string) ([]byte, bool)
	Set(string, []byte)
	Delete(string)
}

type implCache[T any] struct {
	cacheClient CacheClient
	namespace   string
	group       singleflight.Group
	serializer  serializer.Serializer[T]
}

func (c *implCache[T]) key(id string) string {
	return fmt.Sprintf("%s:%s", c.namespace, id)
}

func (c *implCache[T]) Get(
	id string,
	getDataFromPersistentStorage func() (T, error),
) (T, error) {
	var zero T

	// get data from cache
	cachedData, cacheHit := c.cacheClient.Get(c.key(id))
	if cacheHit {
		serializedData, serializeErr := c.serializer.Decode(cachedData)
		if serializeErr == nil {
			return serializedData, nil
		}

		log.Printf("cache data serialization failed for key: %s, error: %v", c.key(id), serializeErr)
		c.Delete(id)
	}

	// get data from persistent storage
	persistentData, persistentErr, _ := c.group.Do(c.key(id), func() (any, error) {
		return getDataFromPersistentStorage()
	})
	if persistentErr != nil {
		return zero, persistentErr
	}

	val, typeOK := persistentData.(T)
	if !typeOK {
		return zero, ErrTypeCastingCacheValueFailed
	}

	c.set(id, val)
	return val, nil
}

func (c *implCache[T]) set(id string, data T) {
	byteData, err := c.serializer.Encode(data)
	if err != nil {
		log.Printf("error serializing cache data for key %s failed: %v", c.key(id), err)
		return
	}

	c.cacheClient.Set(c.key(id), byteData)
}

func (c *implCache[T]) Delete(id string) {
	c.cacheClient.Delete(c.key(id))
}

func NewCache[T any](cacheClient CacheClient, serializer serializer.Serializer[T], namespace string) Cache[T] {
	return &implCache[T]{
		cacheClient: cacheClient,
		namespace:   namespace,
		serializer:  serializer,
	}
}
