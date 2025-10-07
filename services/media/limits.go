package media

const (
	singlepartUploadLimit uint64 = 16 * (1 << 20) // 16 mega byte
	minimumPartSize       uint64 = 5 * (1 << 20)  // 5 mega byte
	maximumPartSize       uint64 = singlepartUploadLimit
	maximumPartsCount     uint16 = 10000
	multipartUploadLimit  uint64 = uint64(maximumPartsCount) * maximumPartSize
	mediaUploadLimit      uint16 = 100
)

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
