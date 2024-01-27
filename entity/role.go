package entity

type Role uint8

const (
	UserRole = 1 + iota
	AdminRole
	SellerRole
)

const (
	UserRoleStr   = "user"
	AdminRoleStr  = "admin"
	SellerRoleStr = "seller"
)

func (r Role) String() string {
	switch r {
	case AdminRole:
		return AdminRoleStr

	case UserRole:
		return UserRoleStr

	case SellerRole:
		return SellerRoleStr
	}
	return ""
}

func MapToRoleEntity(roleStr string) Role {
	switch roleStr {
	case UserRoleStr:
		return UserRole
	case AdminRoleStr:
		return AdminRole

	}
	return Role(0)
}
