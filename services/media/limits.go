package media

const (
	singlepartUploadLimit uint64 = 16 * (1 << 20) // 16 mega byte
	multipartUploadLimit  uint64 = 10 * (1 << 30) // 10 giga byte
	minimumPartSize       uint32 = 5 * (1 << 20)  // 5 mega byte
	maximumPartsCount     uint16 = 10000
	mediaUploadLimit      uint8  = 100
)

type partDetails struct {
	partSize     uint32
	lastPartSize uint32
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

	return &partDetails{}, nil
}
