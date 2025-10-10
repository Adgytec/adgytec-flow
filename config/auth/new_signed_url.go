package auth

import (
	"bytes"
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
	queryKeyActor     = "actor"
)

// NewSignedURL() is using query for cases where other query params are necessary to complete the action
// for majority of request only expire and signature query params are added and rest of the action details are validated using url path and request context
func (a *authCommon) NewSignedURL(actionPath string, query map[string]string, expireAfter time.Duration) (*url.URL, error) {
	if query == nil {
		query = make(map[string]string)
	}

	expire := time.Now().Add(expireAfter).Unix()
	expireString := strconv.FormatInt(expire, 10)

	// add expire details
	query[queryKeyExpire] = expireString

	// remove signature key
	delete(query, queryKeySignature)

	queryKeys := make([]string, 0, len(query))
	for key := range query {
		queryKeys = append(queryKeys, key)
	}
	sort.Strings(queryKeys)

	baseURL := a.apiURL.JoinPath(actionPath)

	var hashPayload bytes.Buffer
	for _, key := range queryKeys {
		hashPayload.WriteString(key)
		hashPayload.WriteString(query[key])
	}

	signedHash, signingErr := a.newSignedHash([]byte(baseURL.Path), hashPayload.Bytes())
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
		// don't add actor to query params
		// actor details are always added using request context
		if key == queryKeyActor {
			continue
		}
		urlQuery.Add(key, query[key])
	}
	baseURL.RawQuery = urlQuery.Encode()

	return baseURL, nil
}

func (a *authCommon) NewSignedURLWithActor(ctx context.Context, actionPath string, query map[string]string, expireAfter time.Duration) (*url.URL, error) {
	actorID, actorErr := actor.GetActorIdFromContext(ctx)
	if actorErr != nil {
		return nil, actorErr
	}

	if query == nil {
		query = make(map[string]string)
	}

	query[queryKeyActor] = actorID.String()
	return a.NewSignedURL(actionPath, query, expireAfter)
}
