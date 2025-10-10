package media

import "github.com/Adgytec/adgytec-flow/database/db"

type uploadInfo struct {
	uploadType   db.GlobalMediaUploadType
	partSize     uint64
	lastPartSize uint64
	partCount    uint16
}

func getUploadInfo(size uint64) (*uploadInfo, error) {
	if size <= singlepartUploadLimit {
		return &uploadInfo{
			uploadType: db.GlobalMediaUploadTypeSinglepart,
		}, nil
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

	return &uploadInfo{
		uploadType:   db.GlobalMediaUploadTypeMultipart,
		partSize:     partSize,
		lastPartSize: lastPartSize,
		partCount:    partCount,
	}, nil
}
