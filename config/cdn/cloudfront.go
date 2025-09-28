package cdn

import (
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/cloudfront/sign"
)

type cdnCloudfront struct {
	urlSigner *sign.URLSigner
	cdnUrl    string
}

func (c *cdnCloudfront) GetSignedUrl(bucketPath *string) *string {
	return nil
}

func NewCloudfrontCDNSigner() (CDN, error) {
	log.Println("creating cloudfront url signer")

	keyPairID := os.Getenv("CLOUDFRONT_KEY_PAIR_ID")
	key := os.Getenv("CLOUDFRONT_PRIVATE_KEY")
	cdnUrl := os.Getenv("CDN_URL")
	if len(keyPairID) == 0 || len(key) == 0 || len(cdnUrl) == 0 {
		return nil, ErrInvalidCloudfrontConfig
	}

	privKey, err := sign.LoadPEMPrivKeyPKCS8AsSigner(strings.NewReader(key))
	if err != nil {
		return nil, &InvalidCloudfrontPrivateKeyError{
			cause: err,
		}
	}

	return &cdnCloudfront{
		urlSigner: sign.NewURLSigner(keyPairID, privKey),
		cdnUrl:    cdnUrl,
	}, nil
}
