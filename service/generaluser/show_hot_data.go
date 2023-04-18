package generaluser

import (
	"fisco/dao/generaluser"

	"github.com/go-redis/redis/v8"
)

func ShowHotData() ([]redis.Z, error) {
	return generaluser.ShowHotData()
}
