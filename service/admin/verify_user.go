package admin

import (
	"fisco/dao/admin"
	"fisco/model"
)

func VerifyUser(userID int, ok bool) (*model.Verify, error) {
	return admin.VerifyUser(userID, ok)

}
