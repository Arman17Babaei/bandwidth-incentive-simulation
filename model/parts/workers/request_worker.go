package workers

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"math/rand"
	"sync"
)

func RequestWorker(requestChan chan types.Request, globalState *types.State, wg *sync.WaitGroup, iterations int) {

	defer wg.Done()
	var requestQueueSize = 10
	var originatorIndex int = 0
	var timeStep int = 0
	var counter int = 0

	defer close(requestChan)

	for counter < iterations {
		if len(requestChan) <= requestQueueSize {

			// TODO: decide on where we should update the timestep. At request creation or request fulfillment
			timeStep = update.Timestep(globalState)
			//timeStep = atomic.LoadInt32(&globalState.TimeStep)
			if timeStep%100 == 0 {
				fmt.Println(timeStep)
			}

			//if timeStep%(iterations/2) == 0 {
			//	fmt.Println("PendingMap is currently:", globalState.PendingStruct.PendingMap)
			//	fmt.Println("RerouteMap is currently:", globalState.RerouteStruct.RerouteMap)
			//}

			originatorIndex = int(update.OriginatorIndex(globalState, timeStep))

			originatorId := globalState.Originators[originatorIndex]

			chunkId := -1
			responsibleNodes := [4]int{}

			if constants.Constants.IsWaitingEnabled() {
				pendingNode := globalState.PendingStruct.GetPending(originatorId)

				if (timeStep-originatorIndex)%constants.Constants.GetEpoke() == 0 || timeStep > iterations {
					if len(pendingNode.ChunkIds) > 0 {
						pendingNode.EpokeDecrement = int32(len(pendingNode.ChunkIds))
						//atomic.AddInt32(&globalState.PendingStruct.Counter, int32(len(pendingNode.ChunkIds)))
					}
				}

				if pendingNode.EpokeDecrement > 0 {
					pendingNodeIds := pendingNode.ChunkIds
					if !globalState.PendingStruct.IsEmpty(originatorId) {
						chunkId = pendingNodeIds[pendingNode.EpokeDecrement-1]
						responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
						pendingNode.EpokeDecrement--
					}

				}
			}

			if constants.Constants.IsRetryWithAnotherPeer() {
				reroute := globalState.RerouteStruct.GetRerouteMap(originatorId)
				if reroute != nil {
					chunkId = reroute[len(reroute)-1]
					responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
				}
			}

			if constants.Constants.IsIterationMeansUniqueChunk() {
				if chunkId == -1 {
					counter++
				}
			} else {
				counter++
			}

			if chunkId == -1 && timeStep < iterations { // No waiting and no retry
				chunkId = rand.Intn(constants.Constants.GetRangeAddress() - 1)

				if constants.Constants.IsPreferredChunksEnabled() {
					var random float32
					numPreferredChunks := 1000
					random = rand.Float32()
					if float32(random) <= 0.5 {
						chunkId = rand.Intn(numPreferredChunks)
					} else {
						chunkId = rand.Intn(constants.Constants.GetRangeAddress()-numPreferredChunks) + numPreferredChunks
					}
				}
				responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
			}

			//if timeStep%(iterations/10) == 0 {
			//	fmt.Println("TimeStep is currently:", timeStep)
			//}
			if chunkId != -1 {
				requestChan <- types.Request{
					OriginatorIndex: originatorIndex,
					OriginatorId:    originatorId,
					TimeStep:        timeStep,
					ChunkId:         chunkId,
					RespNodes:       responsibleNodes,
				}
			}

		}
	}
}
