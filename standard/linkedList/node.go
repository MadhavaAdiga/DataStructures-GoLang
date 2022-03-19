package linkedlist

type node struct {
	data interface{}
	next *node
	prev *node
}

func newNode(data interface{}) *node {
	return &node{
		data: data,
		next: nil,
		prev: nil,
	}
}

// func (node *node) SetNext(ele *node) {
// 	node.next = ele
// }

// func (node *node) SetPrev(ele *node) {
// 	node.prev = ele
// }

// func (node *node) SetData(data interface{}) {
// 	node.data = data
// }
