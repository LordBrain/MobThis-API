package mob

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/LordBrain/MobThis-API/cmd/utils/httperr"
)

func (mob *mobSession) SessionStart() (error httperr.HttpErr) {
	fmt.Println("Session Start now")
	MobSession.Driver = MobSession.Mobbers[rand.Intn(len(MobSession.Mobbers))]
	MobSession.Navigator = MobSession.Mobbers[rand.Intn(len(MobSession.Mobbers))]
	MobSession.State = "Started"

	//Start duration clock
	go func() {
		timer := time.NewTimer(time.Duration(MobSession.Duration) * time.Minute)
		<-timer.C

		fmt.Println("Rotate now")
	}()

	return nil

}
