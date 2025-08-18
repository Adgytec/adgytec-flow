package cache

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type implCache[T any] struct {
	cacheClient core.ICacheClient
	namespace   string
}

func (c *implCache[T]) key(id string) string {
	return fmt.Sprintf("%s:%s", c.namespace, id)
}

func (c *implCache[T]) Get(id string) (T, bool) {
	var zero T

	data, found := c.cacheClient.Get(c.key(id))
	if !found {
		return zero, found
	}

	val, typeOK := data.(T)
	return val, typeOK
}

func (c *implCache[T]) Set(id string, data T) {
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
