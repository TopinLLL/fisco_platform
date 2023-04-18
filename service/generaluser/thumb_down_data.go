package generaluser

import (
	"fisco/dao/generaluser"

	"github.com/gorilla/sessions"
)

func ThumbDownData(session *sessions.Session, dataName, dataContent string) error {
	return generaluser.ThumbDownData(session, dataName, dataContent)
}
