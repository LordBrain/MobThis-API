package mob

import (
	"github.com/LordBrain/MobThis-API/cmd/utils/httperr"
)

func (mob *mobSession) SessionRotate() (error httperr.HttpErr) {

	if MobSession.Driver != "" {
		MobSession.LastDriver = MobSession.Driver
		MobSession.LastNavigator = MobSession.Navigator
		MobSession.Driver = getRandomMobber(MobSession.Mobbers, MobSession.Driver)
		MobSession.Navigator = getRandomMobber(MobSession.Mobbers, MobSession.Driver)

	} else {
		MobSession.Driver = getRandomMobber(MobSession.Mobbers)
		MobSession.Navigator = getRandomMobber(MobSession.Mobbers, MobSession.Driver)
	}

	MobSession.State = "Rotated"

	return nil

}
