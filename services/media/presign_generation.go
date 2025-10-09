package media

import (
	"context"
	"encoding/binary"
	"strconv"
	"time"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/google/uuid"
)

const (
	completeMultipartPresignValidDuration = 4 * time.Hour
)

func (s *mediaService) getUploadDetails(ctx context.Context, newMediaItem NewMediaItemInfoWithStorageDetails) (*MediaUploadDetails, error) {
	itemUploadInfo, itemUploadInfoErr := getUploadInfo(newMediaItem.Size)
	if itemUploadInfoErr != nil {
		return nil, itemUploadInfoErr
	}

	uploadDetails := MediaUploadDetails{
		ID:         newMediaItem.ID,
		UploadType: itemUploadInfo.uploadType,
	}

	// handle singlepart upload
	if itemUploadInfo.uploadType == db.GlobalMediaUploadTypeSinglepart {
		presignPut, presignErr := s.storage.NewPresignPut(ctx, newMediaItem.getMediaItemKey(), newMediaItem.ID)
		if presignErr != nil {
			return nil, presignErr
		}

		uploadDetails.PresignPut = &presignPut
	} else {
		// handle multipart upload
		multipartUploadID, multipartErr := s.storage.NewMultipartUpload(ctx, newMediaItem.getMediaItemKey(), newMediaItem.ID)
		if multipartErr != nil {
			return nil, multipartErr
		}

		// generate presign for each part
		partUploadDetails := make([]MultipartPartUpload, 0, itemUploadInfo.partCount)
		for i := 1; i <= int(itemUploadInfo.partCount); i++ {
			multipartPresignPart, presignErr := s.storage.NewPresignUploadPart(ctx, newMediaItem.getMediaItemKey(), multipartUploadID, int32(i))
			if presignErr != nil {
				return nil, presignErr
			}

			// last part size handling
			partSize := itemUploadInfo.partSize
			if i == int(itemUploadInfo.partCount) {
				partSize = itemUploadInfo.lastPartSize
			}

			part := MultipartPartUpload{
				PartNumber: uint16(i),
				PartSize:   partSize,
				PresignPut: multipartPresignPart,
			}
			partUploadDetails = append(partUploadDetails, part)
		}

		completeMultipartPresign, presignErr := s.newCompleteMultipartPresignGeneration(ctx, newMediaItem.ID)
		if presignErr != nil {
			return nil, presignErr
		}

		uploadDetails.MultipartPresignPart = partUploadDetails
		uploadDetails.MultipartSuccessCallback = &completeMultipartPresign
	}

	return &uploadDetails, nil
}

func (s *mediaService) newCompleteMultipartPresignGeneration(ctx context.Context, mediaID uuid.UUID) (string, error) {
	actorID, actorErr := actor.GetActorIdFromContext(ctx)
	if actorErr != nil {
		return "", actorErr
	}

	expire := time.Now().Add(completeMultipartPresignValidDuration).Unix()

	expireString := strconv.FormatInt(expire, 10)

	expireBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(expireBytes, uint64(expire))

	presignSignature, signatureErr := s.auth.NewSignedHash(actorID[:], []byte(completeMultipartAction), mediaID[:], expireBytes)
	if signatureErr != nil {
		return "", signatureErr
	}

	// create action endpoint
	apiURL := s.apiURL.JoinPath(getCompleteMultipartEndpoint(mediaID))

	// add query params
	urlQuery := apiURL.Query()
	urlQuery.Add("action", completeMultipartAction)
	urlQuery.Add("expires", expireString)
	urlQuery.Add("signature", presignSignature)

	apiURL.RawQuery = urlQuery.Encode()
	return apiURL.String(), nil
}
