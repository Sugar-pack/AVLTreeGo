package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAVLTree_Append(t *testing.T) {
	avlTree := NewAVLTree(1)
	assert.Equal(t, V(1), avlTree.root.value)
	assert.Equal(t, 0, avlTree.BFactor)
	assert.Equal(t, 1, avlTree.Height)

	avlTree.Append(2)
	assert.Equal(t, V(1), avlTree.root.value)
	assert.Equal(t, V(2), avlTree.root.right.root.value)

	avlTree.Append(3)
	assert.Equal(t, V(2), avlTree.root.value)
	assert.Equal(t, V(1), avlTree.root.left.root.value)
	assert.Equal(t, V(3), avlTree.root.right.root.value)
}
