package main

import (
	"fmt"
)

// ======================================================  Item  ===============================================================

type Item struct {
	Val int
}

// ======================================================  Node  ===============================================================

type Node struct {
	id       int
	bucket   *Tree
	items    []*Item
	children []*Node
}

func NewNode(bucket *Tree, items []*Item, children []*Node) *Node {
	return &Node{
		id:       bucket.nextNodeId(),
		bucket:   bucket,
		items:    items,
		children: children,
	}
}

func (n *Node) isOverflowed() bool {
	return len(n.items) > n.bucket.maxItems
}

func (n *Node) isLeaf() bool {
	return n.children == nil || len(n.children) == 0
}

func (n *Node) findKey(val int) (bool, int) {
	for i, item := range n.items {
		if item.Val == val {
			return true, i
		}

		if item.Val > val {
			return false, i
		}
	}

	return false, len(n.items)
}

func (n *Node) add(index int, val int) {
	if len(n.items) == index {
		n.items = append(n.items, &Item{Val: val})
		return
	}

	n.items = append(n.items[:index], append([]*Item{&Item{Val: val}}, n.items[index:]...)...)
}

func (n *Node) split(targetNode *Node, insertionIndex int) {
	i := 0
	nodeSize := n.bucket.minItems

	for targetNode.isOverflowed() {
		midItem := targetNode.items[nodeSize]
		var newNode *Node

		if targetNode.isLeaf() {
			newNode = NewNode(n.bucket, targetNode.items[nodeSize+1:], []*Node{})
			targetNode.items = targetNode.items[:nodeSize]
		} else {
			newNode = NewNode(n.bucket, targetNode.items[nodeSize+1:], targetNode.children[nodeSize+1:])
			targetNode.items = targetNode.items[:nodeSize]
			targetNode.children = targetNode.children[:nodeSize+1]
		}

		n.add(insertionIndex, midItem.Val)

		if len(n.children) == insertionIndex+1 { // if middle list, move items forward
			n.children = append(n.children, newNode)
		} else {
			n.children = append(n.children[:insertionIndex+1], n.children[insertionIndex:]...)
			n.children[insertionIndex+1] = newNode
		}

		insertionIndex += 1
		i++
		targetNode = newNode
	}

}

// ======================================================  Tree  ===============================================================

type Tree struct {
	root       *Node
	minItems   int
	maxItems   int
	nodesCount int
}

func NewTree(minItems int) *Tree {
	return &Tree{
		minItems: minItems,
		maxItems: minItems * 2,
	}
}

func (tree *Tree) nextNodeId() int {
	tree.nodesCount++
	return tree.nodesCount
}

func (tree *Tree) findKey(val int, exact bool) (int, *Node, []int) {
	n := tree.root

	ancestoresIndexes := []int{0} // root index

	for true {
		found, index := n.findKey(val)
		if found {
			return index, n, ancestoresIndexes
		}
		if n.isLeaf() {
			if exact { // for find and delete operations, return nil if not found
				return -1, nil, nil
			}
			return index, n, ancestoresIndexes
		}

		ancestoresIndexes = append(ancestoresIndexes, index)
		n = n.children[index]
	}

	return -1, nil, ancestoresIndexes
}

func (tree *Tree) getNodes(indexes []int) []*Node {
	nodes := []*Node{tree.root}
	child := tree.root
	for i := 1; i < len(indexes); i++ {
		child = child.children[indexes[i]]
		nodes = append(nodes, child)
	}

	return nodes
}

func (tree *Tree) Put(val int) {
	// if root is empty, create new node and insert val
	if tree.root == nil {
		tree.root = &Node{
			id:       tree.nodesCount,
			bucket:   tree,
			items:    []*Item{&Item{Val: val}},
			children: []*Node{},
		}
		tree.nodesCount++
		return
	}

	insertionIndex, nodeToInsert, ancestoresIndexes := tree.findKey(val, false)

	if insertionIndex == -1 {
		panic("Could not find key")
	}

	nodeToInsert.add(insertionIndex, val)

	ancestores := tree.getNodes(ancestoresIndexes)

	// Rebalance the nodes all the way up
	for i := len(ancestoresIndexes) - 2; i >= 0; i-- {
		parentNode := ancestores[i]
		node := ancestores[i+1]
		nodeIndex := ancestoresIndexes[i+1]
		if node.isOverflowed() {
			parentNode.split(node, nodeIndex)
		}
	}

	// Handle root
	if tree.root.isOverflowed() {
		newRoot := NewNode(tree, []*Item{}, []*Node{tree.root})
		newRoot.split(tree.root, 0)
		tree.root = newRoot
	}
}

func (tree *Tree) Find(val int) bool {
	itemIndex, node, _ := tree.findKey(val, true)
	if itemIndex == -1 {
		return false
	}

	return node.items[itemIndex].Val == val
}

// ======================================================  Debug  ===============================================================

func (n *Node) Mermaid() string {
	if n == nil {
		return ""
	}

	output := ""

	// items to comma separated string
	items := fmt.Sprintf("#%d: ", n.id)
	for _, item := range n.items {
		items += fmt.Sprintf("%d, ", item.Val)
	}
	if len(n.items) == 0 {
		items += "EMPTY"
	}

	// loop over children
	for _, child := range n.children {
		output += fmt.Sprintf("%d[%s] --> %d \n", n.id, items, child.id)
		output += child.Mermaid()
	}

	if len(n.children) == 0 {
		output += fmt.Sprintf("%d[%s] \n", n.id, items)
	}

	return output
}

func (tree *Tree) Mermaid() string {
	return tree.root.Mermaid()
}

func (tree *Tree) Print() string {
	return tree.root.Print()
}

func (n *Node) Print() string {
	return n.print(1)
}

func (n *Node) print(indent int) string {
	if n == nil {
		return "\n"
	}

	output := "\n"
	for i := 0; i < indent; i++ {
		output += " "
	}

	itemsString := ""
	for _, item := range n.items {
		itemsString += fmt.Sprintf("%d, ", item.Val)
	}

	output += fmt.Sprintf("Node %d: %s", n.id, itemsString)

	for _, child := range n.children {
		output += child.print(indent * 2)
	}

	return output
}
