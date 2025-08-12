package payload

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

func EncodeJSON[T any](w http.ResponseWriter, status int, data T) {
	jsonRes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Printf("Error encoding json: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(jsonRes)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
}

func EncodeError(w http.ResponseWriter, err error) {
	if responseError, ok := err.(core.IErrorResponse); ok {
		EncodeJSON(w, responseError.HTTPResponse().HTTPStatusCode, responseError.HTTPResponse())
	} else {
		EncodeJSON(w, http.StatusInternalServerError, core.ResponseHTTPError{
			Message: helpers.ValuePtr(http.StatusText(http.StatusInternalServerError)),
		})
	}
}
