package mob

import (
	"github.com/LordBrain/MobThis-API/cmd/utils/httperr"
)

func (mob *mobSession) SessionLeave(mobber string) (error httperr.HttpErr) {

	index := -1

	// Find the index of the target string
	for sliceIndex, mobberName := range MobSession.Mobbers {
		if mobberName == mobber {
			index = sliceIndex
			break
		}
	}

	// If the target string is found, remove it using slicing
	if index != -1 {
		MobSession.Mobbers = append(MobSession.Mobbers[:index], MobSession.Mobbers[index+1:]...)
	}

	return nil

}
