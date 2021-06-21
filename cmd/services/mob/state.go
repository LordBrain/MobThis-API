package mob

import (
	"strconv"
	"time"

	"github.com/LordBrain/MobThis-API/cmd/utils/httperr"
	"github.com/go-redis/redis/v8"
)

func (mob *MobSession) SessionState(redis *redis.Client) (error httperr.HttpErr) {

	mobState, err := redis.HGetAll(CTX, mob.SessionName+"-state").Result()
	if err != nil {
		return httperr.New(500, "Redis Error", err.Error())
	}
	mob.GitRepo = mobState["GitRepo"]
	mob.Driver = mobState["Driver"]
	mob.Navigator = mobState["Navigator"]
	mob.State = mobState["State"]
	mob.Duration, _ = strconv.Atoi(mobState["Duration"])
	switch mobState["Retro"] {
	case "0":
		mob.Retro = false
	case "1":
		mob.Retro = true
	}

	err = redis.LRange(CTX, mob.SessionName+"-mobbers", 0, -1).Err()
	if err != nil {
		return httperr.New(500, "Redis Error", err.Error())
	}

	mob.Mobbers = append(mob.Mobbers, redis.LRange(CTX, mob.SessionName+"-mobbers", 0, -1).Val()...)

	err = redis.Expire(CTX, mob.SessionName+"-state", 60*time.Second).Err()
	if err != nil {
		return httperr.New(500, "Redis Error", err.Error())
	}
	err = redis.Expire(CTX, mob.SessionName+"-mobbers", 60*time.Second).Err()
	if err != nil {
		return httperr.New(500, "Redis Error", err.Error())
	}

	return nil

}
