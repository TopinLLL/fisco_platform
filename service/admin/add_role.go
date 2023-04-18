package admin

import (
	"fisco/dao/admin"
	"fisco/model"
)

func AddRole(roleID int, roleName string) (*model.Role, error) {
	return admin.AddRole(roleID, roleName)
}
