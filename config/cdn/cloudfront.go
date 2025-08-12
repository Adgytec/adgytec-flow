package cdn

import (
	"log"
	"os"
	"strings"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/aws/aws-sdk-go-v2/feature/cloudfront/sign"
)

type cdnCloudfront struct {
	urlSigner *sign.URLSigner
	cdnUrl    string
}

func (c *cdnCloudfront) GetSignedUrl(bucketPath *string) *string {
	return nil
}

func CreateCloudfrontCDNSigner() core.ICDN {
	log.Println("creating cloudfront url signer")

	keyPairId := os.Getenv("CLOUDFRONT_KEY_PAIR_ID")
	key := os.Getenv("CLOUDFRONT_PRIVATE_KEY")
	cdnUrl := os.Getenv("CDN_URL")
	if len(keyPairId) == 0 || len(key) == 0 || len(cdnUrl) == 0 {
		log.Fatalf("Can't find cloudfront key pair id, url and private key")
	}

	privKey, err := sign.LoadPEMPrivKeyPKCS8AsSigner(strings.NewReader(key))
	if err != nil {
		log.Fatalf("Failed to load cloudfront private key, err: %s\n", err.Error())
	}

	return &cdnCloudfront{
		urlSigner: sign.NewURLSigner(keyPairId, privKey),
		cdnUrl:    cdnUrl,
	}
}
