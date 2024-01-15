package mob

import (
	"fmt"
	"time"

	"github.com/LordBrain/MobThis-API/cmd/utils/httperr"
)

func (mob *mobSession) SessionStart() (error httperr.HttpErr) {
	fmt.Println("Session Start now")
	MobSession.State = "Started"

	//Start duration clock
	go func() {
		timer := time.NewTimer(time.Duration(MobSession.Duration) * time.Minute)
		<-timer.C
		// When the duration ends, set the state to rotate so the driver/naviagtor will change
		MobSession.State = "Rotate"
	}()

	return nil

}
