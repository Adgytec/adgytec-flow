package iam

// SelfPermissions defines the permission type used to define self actions
// this is not stored in db
// SelfPermissions are not assignable to any user and are implictly available for all the users for their account actions
type SelfPermissions struct {
	Key  string
	Name string
}
