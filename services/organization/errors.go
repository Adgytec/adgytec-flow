package org

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrMissingRequiredCoreServicesRestrictions = errors.New("missing required core services restrictions")
)

type MissingRequiredCoreServicesRestrictionsError struct {
	missingServicesRestrictions []db.AddServiceRestrictionIntoStagingParams
}

func (e *MissingRequiredCoreServicesRestrictionsError) Error() string {
	var sb strings.Builder
	sb.WriteString("missing required core service restrictions: ")

	for i, r := range e.missingServicesRestrictions {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%s (service: %s)", r.Name, r.ServiceID.String()))
	}

	return sb.String()
}

func (e *MissingRequiredCoreServicesRestrictionsError) Is(target error) bool {
	return target == ErrMissingRequiredCoreServicesRestrictions
}

func (e *MissingRequiredCoreServicesRestrictionsError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}
