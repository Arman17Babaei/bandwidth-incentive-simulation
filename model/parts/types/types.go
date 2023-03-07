package types

type Request struct {
	TimeStep        int32
	OriginatorIndex int32
	OriginatorId    int
	ChunkId         int
	RespNodes       [4]int
}

type Route []int

type Payment struct {
	FirstNodeId  int
	PayNextId    int
	ChunkId      int
	IsOriginator bool
}

type Threshold [2]int

type StateSubset struct {
	OriginatorIndex         int32
	PendingMap              int32
	RerouteMap              int32
	CacheStruct             int32
	SuccessfulFound         int32
	FailedRequestsThreshold int32
	FailedRequestsAccess    int32
	TimeStep                int32
}

type RouteData struct {
	TimeStep int   `json:"timestep"`
	Route    Route `json:"route"`
}

//type StateData struct {
//	TimeStep int         `json:"index"`
//	State    StateSubset `json:"state"`
//}

type State struct {
	Graph                   *Graph
	Originators             []int
	NodesId                 []int
	RouteLists              []Route
	PendingStruct           PendingStruct
	RerouteStruct           RerouteStruct
	CacheStruct             CacheStruct
	OriginatorIndex         int32
	SuccessfulFound         int32
	FailedRequestsThreshold int32
	FailedRequestsAccess    int32
	TimeStep                int32
}

type RequestResult struct {
	Found                bool
	Route                Route
	ThresholdFailedLists [][]Threshold
	AccessFailed         bool
	PaymentList          []Payment
}
