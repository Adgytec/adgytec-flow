package auth

import (
	"context"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/Adgytec/adgytec-flow/utils/actor"
)

const (
	queryKeyExpire    = "expire"
	queryKeySignature = "signature"
)

func (a *authCommon) NewSignedURL(path string, query map[string]string, expireAfter time.Duration) (*url.URL, error) {
	expire := time.Now().Add(expireAfter).Unix()
	expireString := strconv.FormatInt(expire, 10)

	// add expire details
	query[queryKeyExpire] = expireString

	// remove signature key
	delete(query, queryKeySignature)

	queryKeys := make([]string, 0, len(query))
	for key, _ := range query {
		queryKeys = append(queryKeys, key)
	}
	sort.Strings(queryKeys)

	baseURL := *a.apiURL
	baseURL.Path = path

	hashPayload := make([]byte, 0)
	for _, key := range queryKeys {
		hashPayload = append(hashPayload, []byte(query[key])...)
	}

	signedHash, signingErr := a.newSignedHash(hashPayload)
	if signingErr != nil {
		return nil, signingErr
	}

	// add signature to query and slice
	query[queryKeySignature] = signedHash

	// signature should be added at the end of the url
	queryKeys = append(queryKeys, queryKeySignature)

	// add query params
	urlQuery := baseURL.Query()
	for _, key := range queryKeys {
		urlQuery.Add(key, query[key])
	}

	return &baseURL, nil
}

func (a *authCommon) NewSignedURLWithActor(ctx context.Context, path string, query map[string]string, expireAfter time.Duration) (*url.URL, error) {
	actorID, actorErr := actor.GetActorIdFromContext(ctx)
	if actorErr != nil {
		return nil, actorErr
	}

	query["actor"] = actorID.String()
	return a.NewSignedURL(path, query, expireAfter)
}
