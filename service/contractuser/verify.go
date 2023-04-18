package contractuser

import "fisco/dao/contractuser"

func Verify(username, verifyEmail string) error {
	return contractuser.Verify(username, verifyEmail)
}
