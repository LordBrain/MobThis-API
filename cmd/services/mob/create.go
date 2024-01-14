package mob

import (
	"fmt"
	"strings"

	"github.com/LordBrain/MobThis-API/cmd/utils/httperr"
	"github.com/Pallinder/sillyname-go"
)

func (newMob *mobSession) CreateSession(newMobSession NewMobSession) (error httperr.HttpErr) {
	newMobName := strings.ReplaceAll(sillyname.GenerateStupidName(), " ", "-")

	newMob.SessionName = newMobName
	newMob.State = "creating"

	// if len(newMob.Mobbers) == 0 {
	// 	return httperr.New(400, "Missing Mobber", "Missing Mobber")
	// }
	newMob.Mobbers = append(newMob.Mobbers, newMobSession.Mobber)
	newMob.GitRepo = newMobSession.GitRepo
	newMob.Duration = newMobSession.Duration
	newMob.Retro = newMobSession.Retro
	// err := redis.HSet(CTX, newMob.SessionName+"-state", "State", newMob.State, "GitRepo", newMob.GitRepo, "Driver", newMob.Driver, "Navigator", newMob.Navigator, "Duration", newMob.Duration, "Retro", newMob.Retro).Err()
	// if err != nil {
	// 	return httperr.New(500, "Redis Error", err.Error())
	// }
	// err = redis.LPush(CTX, newMob.SessionName+"-mobbers", newMob.Mobbers[0]).Err()
	// if err != nil {
	// 	return httperr.New(500, "Redis Error", err.Error())
	// }

	// err = redis.Expire(CTX, newMob.SessionName+"-state", 60*time.Second).Err()
	// if err != nil {
	// 	return httperr.New(500, "Redis Error", err.Error())
	// }
	// err = redis.Expire(CTX, newMob.SessionName+"-mobbers", 60*time.Second).Err()
	// if err != nil {
	// 	return httperr.New(500, "Redis Error", err.Error())
	// }
	// MobSession = newMob
	fmt.Println(newMob)
	return nil

}
