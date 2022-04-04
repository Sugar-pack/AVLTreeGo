package main

type AVLTree struct {
	root    *Node
	Height  int
	BFactor int
}

type V int

type Node struct {
	value V
	left  *AVLTree
	right *AVLTree
}

func (t *AVLTree) Append(value V) {
	if t.root == nil {
		t.root = &Node{value, nil, nil}
		t.root.left = &AVLTree{}
		t.root.right = &AVLTree{}
	} else if value > t.root.value {
		t.root.left.Append(value)
	} else {
		t.root.right.Append(value)
	}
	t.UpdateHeight()
	t.FixBalance()
}

func (t *AVLTree) UpdateHeight() { //CP
	if t.root != nil {
		leftHeight := 0
		rightHeight := 0
		if t.root.left != nil {
			t.root.left.UpdateHeight()
			leftHeight = t.root.left.Height
		}
		if t.root != nil {
			t.root.right.UpdateHeight()
			rightHeight = t.root.right.Height
		}
		t.Height = max(leftHeight, rightHeight) + 1
	}
	t.Height = 0
}

func (t *AVLTree) UpdateHeightAndBFactor() {
	t.UpdateHeight()
	t.UpdateBFactor()
}

func (t *AVLTree) FixBalance() {
	t.UpdateHeightAndBFactor()
	if t.BFactor == 2 {
		if t.root.left.BFactor < 0 {
			t.root.left.RotateLeft()
			t.UpdateHeightAndBFactor()
			return
		}
		t.RotateRight()
		t.UpdateHeightAndBFactor()
		return
	} else if t.BFactor == -2 {
		if t.root.right.BFactor > 0 {
			t.root.right.RotateRight()
			t.UpdateHeightAndBFactor()
			return
		}
		t.RotateLeft()
		t.UpdateHeightAndBFactor()
		return
	}
}

func (t *AVLTree) UpdateBFactor() {
	if t.root != nil {
		leftHeight := 0
		rightHeight := 0
		if t.root.left != nil {
			t.root.left.UpdateBFactor()
			leftHeight = t.root.left.BFactor
		}
		if t.root != nil {
			t.root.right.UpdateBFactor()
			rightHeight = t.root.right.BFactor
		}
		t.BFactor = rightHeight - leftHeight
	}
	t.BFactor = 0
}

func (t *AVLTree) RotateRight() {
	top := t.root
	mid := t.root.left.root
	bottom := mid.right.root

	t.root = mid
	mid.right.root = top
	top.left.root = bottom
}

func (t *AVLTree) RotateLeft() {
	top := t.root
	mid := t.root.right.root
	bottom := mid.left.root

	t.root = mid
	mid.left.root = top
	top.right.root = bottom
}

func (t *AVLTree) Contains(val V) bool {
	if t.root == nil {
		return false
	}
	if t.root.value == val {
		return true
	}
	if val > t.root.value {
		return t.root.right.Contains(val)
	} else {
		return t.root.left.Contains(val)
	}
}

func NewAVLTree(elems ...V) *AVLTree {
	tree := &AVLTree{
		root: nil,
	}
	for _, elem := range elems {
		tree.Append(elem)
	}
	return tree
}

func max(height int, height2 int) int {
	if height > height2 {
		return height
	}
	return height2
}
