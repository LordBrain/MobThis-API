package mob

import (
	"strings"

	"github.com/LordBrain/MobThis-API/cmd/utils/httperr"
	"github.com/Pallinder/sillyname-go"
)

func (newMob *mobSession) CreateSession(newMobSession NewMobSession) (error httperr.HttpErr) {
	newMobName := strings.ReplaceAll(sillyname.GenerateStupidName(), " ", "-")

	newMob.SessionName = newMobName
	newMob.State = "creating"

	newMob.Mobbers = append(newMob.Mobbers, newMobSession.Mobber)
	newMob.GitRepo = newMobSession.GitRepo
	newMob.Duration = newMobSession.Duration
	newMob.Retro = newMobSession.Retro

	return nil

}
