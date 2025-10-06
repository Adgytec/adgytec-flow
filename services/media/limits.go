package media

const (
	singlepartUploadLimit int = 16 * (1 << 20) // 16 mega byte
	mutipartUploadLimit   int = 10 * (1 << 30) // 10 giga byte
	minimumPartSize       int = 5 * (1 << 20)  // 5 mega byte
	maximumPartsCount     int = 10000
	mediaUploadLimit      int = 100
)

type partDetails struct {
	partSize     int
	lastPartSize int
	partCount    int
}
