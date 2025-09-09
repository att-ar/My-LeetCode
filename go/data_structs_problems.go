package main

/* 706. Design HashMap
89% time, 75% memory

Keys and Values are integers [0, 10^6]
Can do log-structured with a list for fast writes, but I will do an overwriting one for faster reads and no worries of memory usage
At most 10^4 ops will be made on the map

No need for a real hash func
*/

type MyHashMap struct {
	// hashFunc hash.Hash32
	HashMap []*MapNode // public attribute since direct mutations might be useful
	n       int
}

func HashMapConstructor() MyHashMap {
	n := 100
	HashMap := make([]*MapNode, n)
	for i := range n {
		HashMap[i] = &MapNode{}
	}

	return MyHashMap{
		// hashFunc: fnv.New32a(),
		HashMap: HashMap,
		n:       n,
	}
}

func (hm *MyHashMap) getHash(key int) int {
	return key % hm.n
}

func (hm *MyHashMap) Put(key int, value int) {
	root := hm.HashMap[hm.getHash(key)]
	for root.Next != nil {
		root = root.Next
		if root.Key == key {
			root.Val = value
			return
		}
	}
	root.Next = &MapNode{Key: key, Val: value}
}

func (hm *MyHashMap) Get(key int) int {
	root := hm.HashMap[hm.getHash(key)].Next
	for root != nil {
		if root.Key == key {
			return root.Val
		}
		root = root.Next
	}
	return -1 // if not in the map
}

func (hm *MyHashMap) Remove(key int) {
	prev := hm.HashMap[hm.getHash(key)]
	root := prev.Next
	for root != nil {
		if root.Key == key {
			prev.Next = root.Next
			return
		}
		prev = prev.Next
		root = root.Next
	}
}
