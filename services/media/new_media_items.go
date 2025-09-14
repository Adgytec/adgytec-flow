package media

import "github.com/Adgytec/adgytec-flow/utils/core"

func (s *mediaService) newMediaItems(input []NewMediaItemInput) ([]NewMediaItemOutput, error) {
	return nil, core.ErrNotImplemented
}

func (pc *mediaServicePC) NewMediaItems(input []NewMediaItemInput) ([]NewMediaItemOutput, error) {
	return pc.service.newMediaItems(input)
}
