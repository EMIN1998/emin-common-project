package lru

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */

var initLen int
var resourceMap map[int]int
var node *nodeStack

const OperatorAdd = 1
const OperatorGet = 2

func LRU(operators [][]int, k int) []int {
	initLen = k
	initStack()
	resp := make([]int, 0)
	for _, unit := range operators {
		if unit[0] == OperatorAdd {
			set(unit[1], unit[2])
		}

		if unit[0] == OperatorGet {
			tmp := get(unit[1])
			resp = append(resp, tmp)
		}
	}

	return resp
}

type nodeStack struct {
	Len   int
	Deep  int
	stack []int
	// indexMap map[int]int
}

func (n *nodeStack) query(k int) int {
	index := n.getIndex(k)
	if index == -1 {
		return index
	}

	res := n.stack[index]
	n.insertHead(index)

	return resourceMap[res]
}

func (n *nodeStack) getIndex(k int) int {
	var index = -1
	for idx, v := range n.stack {
		if v == k {
			index = idx
			break
		}
	}

	return index
}

func (n *nodeStack) insertHead(index int) {
	list := make([]int, 0)
	list = append(list, n.stack[index])
	n.stack = append(n.stack[:index], n.stack[index+1:]...)
	n.stack = append(list, n.stack...)
}

func (n *nodeStack) insert(k, v int) {
	if n.isFull() {
		delete(resourceMap, k)
		n.stack = n.stack[0 : n.Len-1]
		n.Deep -= 1
	}

	n.Deep += 1
	if _, ok := resourceMap[k]; ok {
		index := n.getIndex(k)
		resourceMap[k] = v
		n.insertHead(index)
		return
	}

	list := make([]int, 0)
	list = append(list, k)
	n.stack = append(list, n.stack...)
	resourceMap[k] = v
	return
}

func (n *nodeStack) isFull() bool {
	if n.Deep < n.Len {
		return false
	}

	return true
}

func initStack() {
	node = &nodeStack{}
	resourceMap = make(map[int]int)
	list := make([]int, initLen)

	node.Len = initLen
	node.stack = list
}

func get(k int) int {
	return node.query(k)
}

func set(k, v int) {
	node.insert(k, v)
}
