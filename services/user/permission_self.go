package user

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
)

var getSelfProfilePermission = core.SelfPermissions{
	Key:  "self:get:user-profile",
	Name: "Get self profile",
}

var updateSelfProfilePermission = core.SelfPermissions{
	Key:  "self:update:user-profile",
	Name: "Update self profile",
}
