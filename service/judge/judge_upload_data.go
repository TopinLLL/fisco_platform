package judge

import "fisco/dao/judge"

func UploadData(agree bool, dataID int) error {
	return judge.UploadData(agree, dataID)
}
