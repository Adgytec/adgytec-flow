package staging

import (
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
)

type Details struct {
	Service                db.AddServicesIntoStagingParams
	ManagementPermissions  []db.AddManagementPermissionsIntoStagingParams
	ApplicationPermissions []db.AddApplicationPermissionsIntoStagingParams
	ServiceRestrictions    []db.AddServiceRestrictionIntoStagingParams
}

type Services interface {
	GetCoreServiceRestrictions() []db.AddServiceRestrictionIntoStagingParams
}

// coreServiceRestrictions will be:
//
// nil if not yet computed
// [] (empty slice) if there are no core restrictions
// populated slice if there are core restrictions
type services struct {
	services                []db.AddServicesIntoStagingParams
	serviceRestrictions     []db.AddServiceRestrictionIntoStagingParams
	coreServiceRestrictions []db.AddServiceRestrictionIntoStagingParams // cached slice
}

func (s *services) GetCoreServiceRestrictions() []db.AddServiceRestrictionIntoStagingParams {
	if s.coreServiceRestrictions != nil {
		return s.coreServiceRestrictions
	}

	coreServiceIDs := make(map[uuid.UUID]struct{})
	for _, svc := range s.services {
		if svc.Type == db.GlobalServiceTypeCore {
			coreServiceIDs[svc.ID] = struct{}{}
		}
	}

	var coreRestrictions []db.AddServiceRestrictionIntoStagingParams
	for _, r := range s.serviceRestrictions {
		if _, ok := coreServiceIDs[r.ServiceID]; ok {
			coreRestrictions = append(coreRestrictions, r)
		}
	}

	s.coreServiceRestrictions = coreRestrictions
	return s.coreServiceRestrictions
}
