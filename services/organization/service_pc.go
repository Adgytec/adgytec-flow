package org

import "github.com/rs/zerolog/log"

type OrgServicePC interface{}

type orgServicePC struct{}

func NewOrgServicePC(params any) OrgServicePC {
	log.Info().
		Str("service", serviceName).
		Msg("new service pc")
	return &orgServicePC{}
}
