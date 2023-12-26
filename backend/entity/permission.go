package entity

import "time"

type Permission struct {
	ID        uint
	Title     PermissionTitle
	CreatedAt time.Time
}
type PermissionTitle string

const (
	UserListPermission     PermissionTitle = PermissionTitle("user-list")
	UserDeletePermission   PermissionTitle = PermissionTitle("user-delete")
	LaptopListPermission   PermissionTitle = PermissionTitle("laptop-list")
	LaptopDeletePermission PermissionTitle = PermissionTitle("laptop-delete")
	LaptopUpdatePermission PermissionTitle = PermissionTitle("laptop-update")
)
