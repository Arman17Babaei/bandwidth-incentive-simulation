package workers

import (
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/update"
)

func UpdateWorker(stateChan chan *State, policyChan chan Policy, globalState *State, stateArray []State, iterations int) {

	for {
		policyOutput := <-policyChan

		UpdatePendingMap(globalState, policyOutput)
		UpdateRerouteMap(globalState, policyOutput)
		UpdateCacheMap(globalState, policyOutput)
		//UpdateOriginatorIndex(globalState)
		UpdateSuccessfulFound(globalState, policyOutput)
		UpdateFailedRequestsThreshold(globalState, policyOutput)
		UpdateFailedRequestsAccess(globalState, policyOutput)
		UpdateRouteListAndFlush(globalState, policyOutput)
		UpdateNetwork(globalState, policyOutput)

		newState := State{
			Graph:                   globalState.Graph,
			Originators:             globalState.Originators,
			NodesId:                 globalState.NodesId,
			RouteLists:              globalState.RouteLists,
			PendingMap:              globalState.PendingMap,
			RerouteMap:              globalState.RerouteMap,
			CacheStruct:             globalState.CacheStruct,
			OriginatorIndex:         globalState.OriginatorIndex,
			SuccessfulFound:         globalState.SuccessfulFound,
			FailedRequestsThreshold: globalState.FailedRequestsThreshold,
			FailedRequestsAccess:    globalState.FailedRequestsAccess,
			TimeStep:                globalState.TimeStep,
		}

		stateArray[newState.TimeStep] = newState
		//fmt.Println(newState.TimeStep)

		//stateChan <- &newState
	}
}
