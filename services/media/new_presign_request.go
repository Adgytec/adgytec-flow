package media

import "github.com/Adgytec/adgytec-flow/utils/core"

func (s *mediaService) newPresignRequest(key string) error {
	return core.ErrNotImplemented
}

func (pc *mediaServicePC) NewPresignRequest(key string) error {
	return pc.service.newPresignRequest(key)
}
