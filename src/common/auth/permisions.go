package auth

const (
	UserCreate = "user:create"
	UserRead   = "user:read"
	UserUpdate = "user:update"
	UserDelete = "user:delete"

	AccountReadOwn = "account:read_own"
	AccountReadAll = "account:read_all"
	AccountBlock   = "account:block"

	TransferCreate = "transfer:create"

	TransactionReadOwn = "transaction:read_own"
	TransactionReadAll = "transaction:read_all"

	RoleRead   = "role:read"
	RoleUpdate = "role:update"
)
