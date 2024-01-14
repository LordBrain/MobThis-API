package mob

import (
	"github.com/LordBrain/MobThis-API/cmd/utils/httperr"
)

func (mob *mobSession) SessionJoin(newUser JoinMob) (error httperr.HttpErr) {

	for _, existingMobber := range MobSession.Mobbers {
		if existingMobber == newUser.Mobber {
			return httperr.New(409, "mobber already exisits", "Mobber has already joined") // Mobber already exists, not appending
		}
	}
	mob.Mobbers = append(MobSession.Mobbers, newUser.Mobber)

	return nil

}
