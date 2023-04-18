package judge

import "fisco/dao/judge"

func JudgeUploadData(agree bool, dataID int) error {
	return judge.JudgeUploadData(agree, dataID)
}
