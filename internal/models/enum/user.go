package enum

type (
	BinaryStatus string

	AccStr string
)

const (
	ACC AccStr = "account"

	// Struct's validation
	UGroup string = "User Group"
	UMail  string = "User's Email"

	// Constanta for status
	ACTIVE   BinaryStatus = "ACTIVE"
	INACTIVE BinaryStatus = "INACTIVE"

	// Role For User
	RoleSuperAdmin string = "SUPER ADMIN"
	RoleAdmin      string = "ADMIN"
	RoleMemeber    string = "MEMBER"
)

var (
	IsValidUserGroup = map[string]bool{
		RoleSuperAdmin: true,
		RoleAdmin:      true,
		RoleMemeber:    true,
	}

	AllAccess = []string{RoleSuperAdmin, RoleAdmin, RoleMemeber}
)
