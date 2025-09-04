package user

import (
	"github.com/Adgytec/adgytec-flow/services/iam"
)

var getSelfProfilePermission = iam.SelfPermissions{
	Key:  "self:get:user-profile",
	Name: "Get self profile",
}

var updateSelfProfilePermission = iam.SelfPermissions{
	Key:  "self:update:user-profile",
	Name: "Update self profile",
}
