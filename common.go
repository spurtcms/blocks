package blocks

import (
	"errors"
	"os"
	"time"
)

var (
	ErrorAuth       = errors.New("auth enabled not initialised")
	ErrorPermission = errors.New("permissions enabled not initialised")
	ErrorEmpty      = errors.New("given some values is empty")
	TZONE, _        = time.LoadLocation(os.Getenv("TIME_ZONE"))
	TenantId        = os.Getenv("Tenant_ID")
)

type BlockModel struct {
	DataAccess int
	UserId     int
}

var Blockmodel BlockModel

func AuthandPermission(block *Block) error {

	//check auth enable if enabled, use auth pkg otherwise it will return error
	if block.AuthEnable && !block.Auth.AuthFlg {

		return ErrorAuth
	}
	//check permission enable if enabled, use team-role pkg otherwise it will return error
	if block.PermissionEnable && !block.Auth.PermissionFlg {

		return ErrorPermission

	}

	return nil
}
