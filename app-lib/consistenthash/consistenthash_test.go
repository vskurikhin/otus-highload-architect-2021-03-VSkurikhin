package consistenthash

import (
	"fmt"
	"testing"
)

var ring = NewRing()

func Test0(t *testing.T) {
	ring.AddNode(0, "db-node-1")
	fmt.Printf("ring.Nodes: %v\n", ring.Nodes)

	name1 := ring.GetHashId("Банжул")
	fmt.Printf("name1: %d\n", name1)

	ring.AddNode(1, "db-node-2")
	fmt.Printf("ring.Nodes: %v\n", ring.Nodes)
	name2 := ring.GetHashId("Банжул")
	fmt.Printf("name2: %d\n", name2)
}

func Test1(t *testing.T) {
	ring.AddNode(0, "db-node-1")
	fmt.Printf("ring.Nodes: %v\n", ring.Nodes)

	name1 := ring.GetId("Банжул")
	fmt.Printf("name1: %d\n", name1)

	ring.AddNode(1, "db-node-2")
	fmt.Printf("ring.Nodes: %v\n", ring.Nodes)
	name2 := ring.GetId("Банжул")
	fmt.Printf("name2: %d\n", name2)
}
