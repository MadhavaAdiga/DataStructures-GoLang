package unionfind

type DisjoinSet struct {
	n          int   // size/capacity of internal array
	rank       []int // rank optimization
	nodes      []int
	components int // no of individual trees
}

func NewDisjoinSet(size int) *DisjoinSet {
	r := make([]int, size)
	n := make([]int, size)

	for i := 0; i < int(size); i++ {
		r[i] = 1
		n[i] = i
	}
	return &DisjoinSet{
		n:          size,
		components: size,
		rank:       r,
		nodes:      n,
	}
}

func (ds *DisjoinSet) Find(k int) int {
	if k >= ds.Size() {
		return -1
	}

	root := k

	for root != ds.nodes[root] {
		root = ds.nodes[root]
	}

	// path compression
	for k != root {
		temp := ds.nodes[k]
		ds.nodes[k] = root
		k = temp
	}

	return root
}

func (ds *DisjoinSet) Union(p, q int) {
	if (p >= ds.Size() || q >= ds.Size()) || (p < 0 || q < 0) {
		return
	}

	if ds.Connected(p, q) {
		return
	}

	root1 := ds.nodes[p]
	root2 := ds.nodes[q]

	// Merge smaller component/set into the larger one
	if ds.rank[root2] > ds.rank[root1] {
		ds.rank[root2] += ds.rank[root1]
		ds.nodes[root1] = root2
		ds.rank[root1] = 0
	} else {
		ds.rank[root1] += ds.rank[root2]
		ds.nodes[root2] = root1
		ds.rank[root2] = 0
	}
	// decrease number of components as number of individual trees is reduced
	ds.components--
}

func (ds *DisjoinSet) Components() int { return ds.components }

func (ds *DisjoinSet) Size() int { return ds.n }

func (ds *DisjoinSet) ComponentSize(k int) int { return ds.rank[ds.Find(k)] }

func (ds *DisjoinSet) Connected(p, q int) bool { return ds.Find(p) == ds.Find(q) }
