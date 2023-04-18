package generaluser

import (
	"fisco/dao/generaluser"

	"github.com/gorilla/sessions"
)

func ThumbUpData(session *sessions.Session, dataName, dataContent string) error {
	return generaluser.ThumbUpData(session, dataName, dataContent)
}
