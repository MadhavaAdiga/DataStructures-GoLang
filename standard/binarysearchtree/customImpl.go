package binarysearchtree

type CustomImpl struct {
	nodeCount int
	root      customeNode
}

func newCustomImpl() BinarySearchTree {
	return &CustomImpl{}
}

func (tree *CustomImpl) Add(interface{}) bool {
	return false
}
func (tree *CustomImpl) Remove(ele interface{}) bool {
	return false
}
func (tree *CustomImpl) IsEmpty() bool {
	return false
}
func (tree *CustomImpl) Contains(interface{}) bool {
	return false
}
func (tree *CustomImpl) Size() int {
	return 0
}
