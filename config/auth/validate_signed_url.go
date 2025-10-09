package auth

import (
	"context"
	"maps"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/Adgytec/adgytec-flow/utils/actor"
)

func validateExpiry(epochString string) error {
	epochInt, parseErr := strconv.ParseInt(epochString, 10, 64)
	if parseErr != nil {
		return &InvalidSignedURLError{}
	}

	expiry := time.Unix(epochInt, 0)
	if time.Now().After(expiry) {
		return &SignedURLExpiredError{}
	}

	return nil
}

// validateSignedURL() baseQuery contains extra query params that are not directly included in url
func (a *authCommon) validateSignedURL(signedURL *url.URL, baseQuery map[string]string) error {
	queryParams := signedURL.Query()

	hashSignatureSlice := queryParams[queryKeySignature]
	expireSlice := queryParams[queryKeyExpire]
	if len(hashSignatureSlice) != 1 && len(expireSlice) != 1 {
		return &InvalidSignedURLError{}
	}

	// remove signature from query params
	delete(queryParams, queryKeySignature)

	hashSignature := hashSignatureSlice[0]
	expireString := expireSlice[0]

	// validate expiry
	expireErr := validateExpiry(expireString)
	if expireErr != nil {
		return expireErr
	}

	// create custom query
	query := make(map[string]string)

	// add query params to query
	for key, value := range queryParams {
		// only consider first element
		query[key] = value[0]
	}

	// add base query to query
	maps.Copy(query, baseQuery)

	// sort query keys
	queryKeys := make([]string, 0, len(query))
	for key := range query {
		queryKeys = append(queryKeys, key)
	}
	sort.Strings(queryKeys)

	hashPayload := make([]byte, 0)
	for _, key := range queryKeys {
		hashPayload = append(hashPayload, []byte(query[key])...)
	}

	// compare hash
	compareErr := a.compareSignedHash(hashSignature, []byte(signedURL.Path), hashPayload)
	if compareErr != nil {
		return compareErr
	}

	return nil
}

func (a *authCommon) ValidateSignedURL(signedURL *url.URL) error {
	return a.validateSignedURL(signedURL, nil)
}

func (a *authCommon) ValidateSignedURLWithActor(ctx context.Context, signedURL *url.URL) error {
	actorID, actorErr := actor.GetActorIdFromContext(ctx)
	if actorErr != nil {
		return actorErr
	}

	// create query and add actor details to it
	query := make(map[string]string)
	query[queryKeyActor] = actorID.String()

	return a.validateSignedURL(signedURL, query)
}
