package dto

type RoleEditList struct {
	RolePermissionID     int    ` json:"rolePermissionID"`
	ModulID              int    ` json:"modulID"`
	RoleID               int    ` json:"roleID"`
	PermissionID         int    ` json:"permissionID"`
	RolePermissionActive int    ` json:"rolePermissionActive"`
	PermissionTitle      string ` json:"permissionTitle"`
	PermissionSlug       string ` json:"permissionSlug"`
}
