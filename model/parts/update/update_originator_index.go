package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

// OriginatorIndex Used by the requestWorker
func OriginatorIndex(state *types.State, timeStep int) int32 {

	curOriginatorIndex := atomic.LoadInt32(&state.OriginatorIndex)
	if constants.GetSameOriginator() {
		if (timeStep)%100 == 0 {
			if int(curOriginatorIndex+1) >= constants.GetOriginators() {
				atomic.StoreInt32(&state.OriginatorIndex, 0)
				return 0
			} else {
				return atomic.AddInt32(&state.OriginatorIndex, 1)
			}
		}
	} else {
		if int(curOriginatorIndex+1) >= constants.GetOriginators() {
			atomic.StoreInt32(&state.OriginatorIndex, 0)
			return 0
		} else {
			if constants.GetSameOriginator() {
				if atomic.LoadInt32(&state.TimeStep)%100 == 0 {
					return atomic.AddInt32(&state.OriginatorIndex, 1)
				}
			} else {
				return atomic.AddInt32(&state.OriginatorIndex, 1)
			}
		}
	}
	return curOriginatorIndex
	//state.OriginatorIndex = rand.Intn(constants.GetOriginators() - 1)
}
