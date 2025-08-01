package payload

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
)

func DecodeRequest[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	// 8 kilo byte
	r.Body = http.MaxBytesReader(w, r.Body, 1<<13)

	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()

	var reqPayload T
	err := jsonDecoder.Decode(&reqPayload)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			message := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return reqPayload, &app_errors.RequestDecodeError{
				Status:  http.StatusBadRequest,
				Message: message,
			}

		case errors.Is(err, io.ErrUnexpectedEOF):
			message := "Request body contains badly-formed JSON"
			return reqPayload, &app_errors.RequestDecodeError{
				Status:  http.StatusBadRequest,
				Message: message,
			}

		case errors.As(err, &unmarshalTypeError):
			message := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return reqPayload, &app_errors.RequestDecodeError{
				Status:  http.StatusBadRequest,
				Message: message,
			}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			message := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return reqPayload, &app_errors.RequestDecodeError{
				Status:  http.StatusBadRequest,
				Message: message,
			}

		case errors.Is(err, io.EOF):
			message := "Request body must not be empty"
			return reqPayload, &app_errors.RequestDecodeError{
				Status:  http.StatusBadRequest,
				Message: message,
			}

		case err.Error() == "http: request body too large":
			message := "Request body must not be larger than 8 kilobyte."
			return reqPayload, &app_errors.RequestDecodeError{
				Status:  http.StatusRequestEntityTooLarge,
				Message: message,
			}

		default:
			log.Printf("Error decoding request body: %v\n", err)
			return reqPayload, err
		}
	}

	err = jsonDecoder.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		message := "Request body must only contain a single JSON object"
		return reqPayload, &app_errors.RequestDecodeError{
			Status:  http.StatusBadRequest,
			Message: message,
		}
	}

	return reqPayload, nil
}
