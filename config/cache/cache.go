package cache

import (
	"fmt"
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"golang.org/x/sync/singleflight"
)

type implCache[T any] struct {
	cacheClient core.ICacheClient
	namespace   string
	group       singleflight.Group
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
		val, typeOK := cachedData.(T)
		if typeOK {
			return val, nil
		}

		log.Printf("cache type-casting failed for key: %s", c.key(id))
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
		return zero, app_errors.ErrTypeCastingCacheValueFailed
	}

	c.set(id, val)
	return val, nil
}

func (c *implCache[T]) set(id string, data T) {
	c.cacheClient.Set(c.key(id), data)
}

func (c *implCache[T]) Delete(id string) {
	c.cacheClient.Delete(c.key(id))
}

func CreateNewCache[T any](cacheClient core.ICacheClient, namespace string) core.ICache[T] {
	return &implCache[T]{
		cacheClient: cacheClient,
		namespace:   namespace,
	}
}
