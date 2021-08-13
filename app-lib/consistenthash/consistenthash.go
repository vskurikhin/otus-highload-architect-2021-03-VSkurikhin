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
	Id     uint8
	Name   string
	HashId uint8
}

func NewRing() *Ring {
	return &Ring{Nodes: Nodes{}}
}

func NewNode(id uint8, name string) *Node {
	return &Node{
		Id:     id,
		Name:   name,
		HashId: uint8(crc32.Checksum([]byte(name), crc32.MakeTable(2)) % 128),
	}
}

func (r *Ring) AddNode(id uint8, name string) {
	node := NewNode(id, name)
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

func (r *Ring) GetHashId(key string) uint8 {
	return r.get(key).HashId
}

func (r *Ring) GetId(key string) uint8 {
	return r.get(key).Id
}

func (r *Ring) GetName(key string) string {
	return r.get(key).Name
}

func (r *Ring) get(key string) Node {
	searchfn := func(i int) bool {
		return r.Nodes[i].HashId >= uint8(crc32.Checksum([]byte(key), crc32.MakeTable(2))%128)
	}
	i := sort.Search(r.Nodes.Len(), searchfn)
	if i >= r.Nodes.Len() {
		i = 0
	}
	return r.Nodes[i]
}
