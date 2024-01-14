package mob

import (
	"github.com/LordBrain/MobThis-API/cmd/utils/httperr"
)

func (mob *mobSession) SessionLeave() (error httperr.HttpErr) {

	// if len(mob.Mobbers) == 0 {
	// 	return httperr.New(400, "Missing Mobber", "Missing Mobber")
	// }
	// err := redis.LRem(CTX, mob.SessionName+"-mobbers", 1, mob.Mobbers[0]).Err()
	// if err != nil {
	// 	return httperr.New(500, "Redis Error", err.Error())
	// }
	// err = redis.Expire(CTX, mob.SessionName+"-mobbers", 60*time.Second).Err()
	// if err != nil {
	// 	return httperr.New(500, "Redis Error", err.Error())
	// }

	return nil

}
