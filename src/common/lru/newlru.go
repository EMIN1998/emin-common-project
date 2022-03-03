package lru

type LinkListNode struct {
	next *LinkListNode
	prev *LinkListNode
	val  int
	key  int
}

func initLinkList(newKey, value int) *LinkListNode {
	node := &LinkListNode{
		val:  value,
		key:  newKey,
		next: nil,
		prev: nil,
	}

	return node
}

type LRUCache struct {
	valueMap   map[int]*LinkListNode
	cap        int
	size       int
	head, tail *LinkListNode
}

func (s *LRUCache) addNewNode(newNode *LinkListNode) {
	tmp := s.head.next
	newNode.next = tmp
	tmp.prev = newNode
	s.head.next = newNode
	newNode.prev = s.head
}

func (s *LRUCache) removeNode(oldNode *LinkListNode) {
	// pre := oldNode.pre
	oldNode.next.prev = oldNode.prev
	oldNode.prev.next = oldNode.next

	oldNode.next = nil
	oldNode.prev = nil
}

func (s *LRUCache) removeTail() *LinkListNode {
	node := s.tail.prev
	s.removeNode(node)

	return node
}

func (s *LRUCache) moveToHead(node *LinkListNode) {
	s.removeNode(node)
	s.addNewNode(node)
}

func (s *LRUCache) CPush(key, value int) {
	if _, ok := s.valueMap[key]; !ok {
		node := &LinkListNode{
			val: value,
			key: key,
		}

		s.valueMap[key] = node
		s.addNewNode(node)
		s.size++

		if s.size > s.cap {
			node = s.removeTail()
			delete(s.valueMap, node.key)
			s.size--
		}

		return
	}

	node := s.valueMap[key]
	node.val = value

	s.moveToHead(node)
}

func (s *LRUCache) CPop(key int) int {
	if _, ok := s.valueMap[key]; !ok {
		return -1
	}

	node := s.valueMap[key]
	s.moveToHead(node)
	return node.val
}

func Constructor(capacity int) LRUCache {
	// var headNode, tailNode *LinkListNode
	lru := LRUCache{
		head:     initLinkList(0, 0),
		tail:     initLinkList(0, 0),
		size:     0,
		cap:      capacity,
		valueMap: make(map[int]*LinkListNode),
	}

	lru.head.next = lru.tail
	lru.tail.prev = lru.head

	return lru
}

func (this *LRUCache) Get(key int) int {
	return this.CPop(key)
}

func (this *LRUCache) Put(key int, value int) {
	this.CPush(key, value)
}
