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
	missing := make([]string, len(e.missingServicesRestrictions))
	for i, r := range e.missingServicesRestrictions {
		missing[i] = fmt.Sprintf("%s (service: %s)", r.Name, r.ServiceID.String())
	}

	return "missing required core service restrictions: " + strings.Join(missing, ", ")
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
