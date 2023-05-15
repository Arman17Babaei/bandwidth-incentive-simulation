package output

import (
	"bufio"
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/utils"
	"os"
	"sort"
)

type WorkInfo struct {
	ForwardMap map[int]int
	WorkMap    map[int]int
	Requests   map[int]int
	File       *os.File
	Writer     *bufio.Writer
}

func InitWorkInfo() *WorkInfo {
	winfo := WorkInfo{}
	winfo.ForwardMap = make(map[int]int)
	winfo.WorkMap = make(map[int]int)
	winfo.Requests = make(map[int]int)
	winfo.File = MakeFile("./results/work.txt")
	winfo.Writer = bufio.NewWriter(winfo.File)
	LogExpSting(winfo.Writer)
	return &winfo
}

func (wi *WorkInfo) Close() {
	err := wi.Writer.Flush()
	if err != nil {
		fmt.Println("Couldn't flush the remaining buffer in the writer for work output")
	}
	err = wi.File.Close()
	if err != nil {
		fmt.Println("Couldn't close the file with filepath: ./results/work.txt")
	}
}

func (o *WorkInfo) CalculateWorkFairness() float64 {
	size := config.GetNetworkSize()
	vals := make([]int, size)
	i := 0
	for _, value := range o.WorkMap {
		vals[i] = value
		i++
	}
	return utils.Gini(vals)
}

func (o *WorkInfo) CalculateForwardWorkFairness() float64 {
	size := config.GetNetworkSize()
	vals := make([]int, size)
	i := 0
	for _, value := range o.ForwardMap {
		vals[i] = value
		i++
	}
	return utils.Gini(vals)
}

// calculate the maximum work done,
// maximum work done by not originator and
// median work done.
func (o *WorkInfo) CalculateMaxMedianWork() (int, int, int) {
	vals := make([]int, 0, len(o.WorkMap))

	maxfwd := 0

	for id, value := range o.ForwardMap {
		vals = append(vals, value)
		if value > maxfwd && o.Requests[id] == 0 {
			maxfwd = value
		}
	}
	sort.Slice(vals, func(i2, j int) bool {
		return vals[i2] < vals[j]
	})

	return vals[len(vals)-1], maxfwd, vals[len(vals)/2]
}

func (wi *WorkInfo) Update(output *types.OutputStruct) {
	route := output.RouteWithPrices
	for i, hop := range route {
		requester := int(hop.RequesterNode)
		provider := int(hop.ProviderNode)
		if i == 0 {
			wi.Requests[requester]++
		}

		if i != len(route)-1 {
			wi.ForwardMap[provider]++
		}
		wi.WorkMap[provider]++
	}
}

func (wi *WorkInfo) Log() {
	workFairness := wi.CalculateWorkFairness()
	forwardFairness := wi.CalculateForwardWorkFairness()
	max, maxfwd, median := wi.CalculateMaxMedianWork()
	_, err := wi.Writer.WriteString(fmt.Sprintf("Workfairness: %f  \n", workFairness))
	if err != nil {
		panic(err)
	}
	_, err = wi.Writer.WriteString(fmt.Sprintf("Forwardworkfairness: %f  \n", forwardFairness))
	if err != nil {
		panic(err)
	}
	_, err = wi.Writer.WriteString(fmt.Sprintf("Max, max by non originator, and median work done: %d, %d, %d \n", max, maxfwd, median))
	if err != nil {
		panic(err)
	}
}
