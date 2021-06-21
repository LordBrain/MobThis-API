package mob

import (
	"strings"
	"time"

	"github.com/LordBrain/MobThis-API/cmd/utils/httperr"
	"github.com/Pallinder/sillyname-go"
	"github.com/go-redis/redis/v8"
)

func (newMob *MobSession) CreateSession(redis *redis.Client) (error httperr.HttpErr) {
	newMobName := strings.ReplaceAll(sillyname.GenerateStupidName(), " ", "-")

	newMob.SessionName = newMobName
	newMob.State = "creating"

	if len(newMob.Mobbers) == 0 {
		return httperr.New(400, "Missing Mobber", "Missing Mobber")
	}
	err := redis.HSet(CTX, newMob.SessionName+"-state", "State", newMob.State, "GitRepo", newMob.GitRepo, "Driver", newMob.Driver, "Navigator", newMob.Navigator, "Duration", newMob.Duration, "Retro", newMob.Retro).Err()
	if err != nil {
		return httperr.New(500, "Redis Error", err.Error())
	}
	err = redis.LPush(CTX, newMob.SessionName+"-mobbers", newMob.Mobbers[0]).Err()
	if err != nil {
		return httperr.New(500, "Redis Error", err.Error())
	}


	err = redis.Expire(CTX, newMob.SessionName+"-state", 60*time.Second).Err()
	if err != nil {
		return httperr.New(500, "Redis Error", err.Error())
	}
	err = redis.Expire(CTX, newMob.SessionName+"-mobbers", 60*time.Second).Err()
	if err != nil {
		return httperr.New(500, "Redis Error", err.Error())
	}

	return nil

}
