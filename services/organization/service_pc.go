package org

import "github.com/rs/zerolog/log"

type OrgServicePC interface{}

type orgServicePC struct {
	service *orgService
}

func NewOrgServicePC(params orgServiceParams) OrgServicePC {
	log.Info().
		Str("service", serviceName).
		Msg("new service pc")
	return &orgServicePC{
		service: newOrgService(params),
	}
}
