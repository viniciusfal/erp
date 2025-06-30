package rbac

var RolePermissions = map[string][]string{
	"admin": {
		"accountability.approve",
		"accountability.reject",
		"accountability.view_requests",
		"accountability.create",
		"accountability.update",
		"accountability.view",
	},
	"correntista": {
		"accountability.create",
		"accountability.view_self",
	},
	"manager": {
		"accountability.view_requests",
		"accountability.create",
	},
}
