package media

type partDetails struct {
	partSize     uint64
	lastPartSize uint64
	partCount    uint16
}

func getPartDetails(size uint64) (*partDetails, error) {
	if size < singlepartUploadLimit {
		return nil, ErrMultipartTooSmall
	}

	if size > multipartUploadLimit {
		return nil, &MediaTooLargeError{
			Size: size,
		}
	}

	// Start with minimum part size
	partSize := minimumPartSize
	partCount := uint16((size + partSize - 1) / partSize) // ceil division

	// Increase part size if partCount exceeds maximumPartsCount
	if partCount > maximumPartsCount {
		partSize = (size + uint64(maximumPartsCount) - 1) / uint64(maximumPartsCount)
		partCount = uint16((size + partSize - 1) / partSize)
	}

	lastPartSize := partSize
	if size%partSize != 0 {
		lastPartSize = size % partSize
	}

	return &partDetails{
		partSize:     partSize,
		lastPartSize: lastPartSize,
		partCount:    partCount,
	}, nil
}
