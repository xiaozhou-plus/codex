package bst

type TreeNode struct {
	Key   int
	Value int

	N int

	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Get(key int) int {
	if t == nil {
		return -1
	}

	if key < t.Key {
		return t.Left.Get(key)
	} else if key > t.Key {
		return t.Right.Get(key)
	} else {
		return t.Value
	}
}

func (t *TreeNode) Size() int {
	if t == nil {
		return 0
	}

	return t.N
}
