package media

func (s *mediaService) newMediaItems(input []NewMediaItemInput) ([]NewMediaItemOutput, error) {
	return nil, nil
}

func (pc *mediaServicePC) NewMediaItems(input []NewMediaItemInput) ([]NewMediaItemOutput, error) {
	return pc.service.newMediaItems(input)
}
