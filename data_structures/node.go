package linkedlist

// base struct
type _node struct {
	data interface{}
}

/*
node is a singly linked list node
*/
type node struct {
	_node //embeded struct
	next  *node
}

/*
doublyNode is a doubly linked list node
*/
type doublyNode struct {
	_node
	next *doublyNode
	prev *doublyNode
}
