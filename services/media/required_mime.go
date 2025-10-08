package media

var (
	AllMime    = []string{"application/octet-stream"}
	ImageMime  = []string{"image/"}
	VideoMime  = []string{"video/"}
	VisualMime = append(ImageMime, VideoMime...)
)
