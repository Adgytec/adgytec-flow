package payload

import (
	"encoding/json"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	"github.com/rs/zerolog/log"
)

func EncodeJSON[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.SetIndent("", "\t")

	if err := jsonEncoder.Encode(data); err != nil {
		log.Error().
			Err(err).
			Str("action", "encode json").
			Send()
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
}

func EncodeError(w http.ResponseWriter, err error) {
	if responseError, ok := err.(apires.ErrorResponse); ok {
		EncodeJSON(w, responseError.HTTPResponse().HTTPStatusCode, responseError.HTTPResponse())
	} else {
		EncodeJSON(w, http.StatusInternalServerError, apires.ErrorDetails{
			Message: pointer.New(http.StatusText(http.StatusInternalServerError)),
		})
	}
}
