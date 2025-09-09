package media

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var managementPermissions = []db.AddManagementPermissionParams{
	completeMediaPipelinePermission,
}

var completeMediaPipelinePermission = db.AddManagementPermissionParams{
	Key:       fmt.Sprintf("%s:pipeline:complete", mediaServiceDetails.Name),
	ServiceID: mediaServiceDetails.ID,
	Name:      "Complete Media Pipeline",
	Description: pointer.New(`
### Complete Media Pipeline

Grants the ability to trigger the completion of a media processing pipeline for a media object.
`),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeApiKey,
}
