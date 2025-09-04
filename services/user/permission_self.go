package user

import (
	"github.com/Adgytec/adgytec-flow/database/models"
)

var getSelfProfilePermission = models.SelfPermissions{
	Key:  "self:get:user-profile",
	Name: "Get self profile",
}

var updateSelfProfilePermission = models.SelfPermissions{
	Key:  "self:update:user-profile",
	Name: "Update self profile",
}
