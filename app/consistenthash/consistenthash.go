package consistenthash

import (
	"hash/crc32"
	"sort"
)

// Ring это сеть распределенных узлов.
type Ring struct {
	Nodes Nodes
}

// Nodes это массив узлов.
type Nodes []Node

// Node это единое целое в кольце.
type Node struct {
	Id     string
	HashId uint8
}

func NewRing() *Ring {
	return &Ring{Nodes: Nodes{}}
}
func NewNode(id string) *Node {
	return &Node{
		Id:     id,
		HashId: uint8(crc32.Checksum([]byte(id), crc32.MakeTable(0)) % 255),
	}
}

func (r *Ring) AddNode(id string) {
	node := NewNode(id)
	r.Nodes = append(r.Nodes, *node)
	sort.Sort(r.Nodes)
}

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Less(i, j int) bool {
	return n[i].HashId < n[j].HashId
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (r *Ring) Get(key string) string {
	searchfn := func(i int) bool {
		return r.Nodes[i].HashId >= uint8(crc32.Checksum([]byte(key), crc32.MakeTable(0)) % 255)
	}
	i := sort.Search(r.Nodes.Len(), searchfn)
	if i >= r.Nodes.Len() {
		i = 0
	}
	return r.Nodes[i].Id
}