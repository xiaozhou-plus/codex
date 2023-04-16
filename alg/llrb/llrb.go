package llrb

type Item interface {
	Less(than Item) bool
}

func less(a, b Item) bool {
	if a == pinf {
		return false
	}

	if a == ninf {
		return true
	}

	return a.Less(b)
}

var (
	pinf = pInf{}
	ninf = nInf{}
)

type pInf struct{}

func (pInf) Less(than Item) bool {
	return true
}

type nInf struct{}

func (nInf) Less(than Item) bool {
	return false
}

type Node struct {
	Item

	Left, Right *Node
	Black       bool
}

type LLRB struct {
	root  *Node
	count int
}

func New() *LLRB {
	return &LLRB{}
}

func (l *LLRB) SetRoot(root *Node) {
	l.root = root
}

func (l *LLRB) GetRoot() *Node {
	return l.root
}

func (l *LLRB) GetCount() int {
	return l.count
}

func (l *LLRB) Min() Item {
	if l.root == nil {
		return nil
	}

	p := l.root
	for p.Left != nil {
		p = p.Left
	}

	return p.Item
}

func (l *LLRB) Max() Item {
	if l.root == nil {
		return nil
	}

	p := l.root
	for p.Right != nil {
		p = p.Right
	}

	return p.Item
}

func (l *LLRB) Get(key Item) Item {
	p := l.root
	for p != nil {
		switch {
		case less(key, p.Item):
			p = p.Left
		case less(p.Item, key):
			p = p.Right
		default:
			return p.Item
		}
	}
	return nil
}

func (t *LLRB) ReplaceOrInsertBulk(items ...Item) {
	for _, i := range items {
		t.ReplaceOrInsert(i)
	}
}

func (t *LLRB) ReplaceOrInsert(item Item) Item {
	if item == nil {
		panic("inserting nil item")
	}
	var replaced Item
	t.root, replaced = t.replaceOrInsert(t.root, item)
	t.root.Black = true
	if replaced == nil {
		t.count++
	}
	return replaced
}

func (t *LLRB) replaceOrInsert(h *Node, item Item) (*Node, Item) {
	if h == nil {
		return newNode(item), nil
	}

	var replaced Item
	switch {
	case less(item, h.Item):
		h.Left, replaced = t.replaceOrInsert(h.Left, item)
	case less(h.Item, item):
		h.Right, replaced = t.replaceOrInsert(h.Right, item)
	default:
		replaced, h.Item = h.Item, item
	}

	walkUpRot23(h)

	return h, replaced
}

func (l *LLRB) DeleteMin() {

}

// Internal node manipulation routines
func newNode(item Item) *Node {
	return &Node{Item: item}
}

func rotateLeft(h *Node) *Node {
	x := h.Right
	h.Right = x.Left

	x.Left = h

	x.Black = h.Black
	h.Black = false
	return x
}

func rotateRight(h *Node) *Node {
	x := h.Left
	h.Left = x.Right
	x.Right = h

	x.Black = h.Black
	h.Black = false
	return x
}

func flip(h *Node) {
	h.Black = !h.Black
	h.Left.Black = !h.Left.Black
	h.Right.Black = !h.Right.Black
}

func walkUpRot23(h *Node) *Node {
	if isRed(h.Right) && !isRed(h.Left) {
		h = rotateLeft(h)
	}

	if isRed(h.Left) && isRed(h.Left.Left) {
		h = rotateRight(h)
	}

	if isRed(h.Left) && isRed(h.Right) {
		flip(h)
	}

	return h
}

func isRed(h *Node) bool {
	if h == nil {
		return false
	}

	return !h.Black
}
