package mob

import (
	"github.com/LordBrain/MobThis-API/cmd/utils/httperr"
)

// SessionEnd will reset the mob session to a empty state
func (mob *mobSession) SessionEnd() (error httperr.HttpErr) {

	MobSession = mobSession{}

	return nil

}
