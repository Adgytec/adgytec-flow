package cdn

import (
	"os"
	"path"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/cloudfront/sign"
	"github.com/rs/zerolog/log"
)

type cdnCloudfront struct {
	urlSigner *sign.URLSigner
	cdnUrl    string
}

func (c *cdnCloudfront) generateURL(bucketPath string) string {
	return path.Join(c.cdnUrl, bucketPath)
}

func (c *cdnCloudfront) GetSignedUrl(bucketPath *string) *string {
	if bucketPath == nil {
		return nil
	}

	signedURL, signErr := c.urlSigner.Sign(c.generateURL(*bucketPath),
		time.Now().Add(signedURLExpiration),
	)
	if signErr != nil {
		log.Error().Err(signErr).Str("action", "get-signed-url").Str("bucket-path", *bucketPath).Send()
		return nil
	}

	return &signedURL
}

func NewCloudfrontCDNSigner() (CDN, error) {
	log.Info().Msg("new cloudfront cdn signer")
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
