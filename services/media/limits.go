package media

const (
	singlepartUploadLimit uint64 = 16 * (1 << 20) // 16 mega byte
	minimumPartSize       uint64 = 5 * (1 << 20)  // 5 mega byte
	maximumPartSize       uint64 = singlepartUploadLimit
	minimumPartsCount     uint16 = 4
	maximumPartsCount     uint16 = 10000
	multipartUploadLimit  uint64 = uint64(maximumPartsCount) * maximumPartSize
	mediaUploadLimit      uint16 = 100
)
