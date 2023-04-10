package bst

type BinarySearchTree struct {
	Root *TreeNode
}

func (b *BinarySearchTree) Get(key int) int {
	return get(b.Root, key)
}

func get(root *TreeNode, key int) int {
	if root == nil {
		return -1
	}

	if key < root.Key {
		return get(root.Left, key)
	} else if key > root.Key {
		return get(root.Right, key)
	} else {
		return root.Value
	}
}
func (b *BinarySearchTree) Put(key, value int) {
	b.Root = put(b.Root, key, value)
	return
}

func put(root *TreeNode, key, value int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Key:   key,
			Value: value,
			N:     1,
		}
	}

	if key < root.Key {
		root.Left = put(root.Left, key, value)
	} else if key > root.Value {
		root.Right = put(root.Right, key, value)
	} else {
		root.Value = value
	}
	root.N = root.Left.Size() + root.Right.Size() + 1
	return root
}

func (b *BinarySearchTree) Min() int {
	return minNode(b.Root).Key
}

func minNode(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}

	return minNode(root.Left)
}

func (b *BinarySearchTree) Max() int {
	return maxNode(b.Root).Value
}

func maxNode(root *TreeNode) *TreeNode {
	if root.Right == nil {
		return root
	}

	return maxNode(root.Right)
}
