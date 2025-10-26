package staging

import (
	"sync"

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
	initOnce                sync.Once
}

func (s *services) GetCoreServiceRestrictions() []db.AddServiceRestrictionIntoStagingParams {
	s.initOnce.Do(func() {
		coreServiceIDs := make(map[uuid.UUID]struct{})
		for _, svc := range s.services {
			if svc.Type == db.GlobalServiceTypeCore {
				coreServiceIDs[svc.ID] = struct{}{}
			}
		}

		coreRestrictions := make([]db.AddServiceRestrictionIntoStagingParams, 0)
		for _, r := range s.serviceRestrictions {
			if _, ok := coreServiceIDs[r.ServiceID]; ok {
				coreRestrictions = append(coreRestrictions, r)
			}
		}

		s.coreServiceRestrictions = coreRestrictions
	})

	return s.coreServiceRestrictions
}

func NewServices(details []Details) Services {
	svcList := make([]db.AddServicesIntoStagingParams, 0, len(details))
	restrictionList := make([]db.AddServiceRestrictionIntoStagingParams, 0)

	for _, d := range details {
		svcList = append(svcList, d.Service)
		if len(d.ServiceRestrictions) > 0 {
			restrictionList = append(restrictionList, d.ServiceRestrictions...)
		}
	}

	return &services{
		services:            svcList,
		serviceRestrictions: restrictionList,
		// coreServiceRestrictions left nil for lazy caching
	}
}
