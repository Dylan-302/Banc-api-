package auth

var RolePermissions = map[string][]string{

	"admin": {
		UserCreate,
		UserRead,
		UserUpdate,
		UserDelete,
		RoleRead,
		RoleUpdate,
	},

	"customer": {
		TransferCreate,
	},

	"employee": {
		UserRead,
	},
}
