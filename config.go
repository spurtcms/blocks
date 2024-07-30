package blocks

import (
	"github.com/spurtcms/auth"
	role "github.com/spurtcms/team-roles"
	"gorm.io/gorm"
)

type Config struct {
	AuthEnable       bool
	PermissionEnable bool
	DB               *gorm.DB
	Auth             *auth.Auth
	Permissions      *role.PermissionConfig
}

type Block struct {
	AuthEnable       bool
	PermissionEnable bool
	AuthFlg          bool
	PermissionFlg    bool
	DB               *gorm.DB
	Auth             *auth.Auth
	Permissions      *role.PermissionConfig
	DataAccess       int
	UserId           int
}
