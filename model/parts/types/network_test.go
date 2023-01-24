package types

import (
	"testing"
)

func TestNetwork(t *testing.T) {
	path := "../../../data/nodes_data_8_10000.txt"
	network := Network{}
	bits, bin, nodes := network.Load(path)

	t.Log("Bits:", bits)
	t.Log("Bin:", bin)
	//print the Nodes map
	for k, v := range nodes {
		t.Log("Nodes:", k, *v)
	}

	t.Log("Nodes[12381]:", *nodes[12381])

	for _, bucket := range nodes[12381].Adj {
		for _, node := range bucket {
			t.Log("Nodes[12381].adj:", node.Id)
		}
		t.Log("\n")
	}
}

// func TestGetNodeById(t *testing.T) {

// }
