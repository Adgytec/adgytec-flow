package org

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/markdown"

	md "github.com/nao1215/markdown"
)

var managementPermissions = []db.AddManagementPermissionsIntoStagingParams{
	createOrganizationPermission,
}

var createOrganizationPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:create:organization", orgServiceDetails.Name),
	ServiceID: orgServiceDetails.ID,
	Name:      "Create Organization",

	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Create Organization").
			PlainText("Grants the ability to create new organization.")
	}),
	RequiredResources: nil,
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}
