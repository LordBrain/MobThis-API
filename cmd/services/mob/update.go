package mob

import (
	"github.com/LordBrain/MobThis-API/cmd/utils/httperr"
	"github.com/go-redis/redis/v8"
)

func (updateMob *MobSession) UpdateSession(redis *redis.Client) (error httperr.HttpErr) {

	//TODO Catch empty list
	err := redis.HSet(CTX, updateMob.SessionName+"-state", "State", updateMob.State, "GitRepo", updateMob.GitRepo, "Driver", updateMob.Driver, "Navigator", updateMob.Navigator, "Duration", updateMob.Duration, "Retro", updateMob.Retro).Err()
	if err != nil {
		return httperr.New(500, "Redis Error", err.Error())
	}

	return nil

}
