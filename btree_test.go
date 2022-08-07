package main

import "testing"

func Test_tree_initialized_with_empty_items_and_children(t *testing.T) {
	tree := NewTree(2)

	if tree.minItems != 2 {
		t.Errorf("Expected tree.minItems to be 2, got %d", tree.minItems)
	}
	if tree.maxItems != 4 {
		t.Errorf("Expected tree.maxItems to be 4, got %d", tree.maxItems)
	}
}

func Test_tree_inserts_item_into_root_node(t *testing.T) {
	tree := NewTree(2)
	tree.Put(1)

	if len(tree.root.items) != 1 {
		t.Errorf("Expected root.items to have 1 item, got %d", len(tree.root.items))
	}
	if tree.root.items[0].Val != 1 {
		t.Errorf("Expected root.items[0].Val to be 1, got %d", tree.root.items[0].Val)
	}
}

func Test_tree_inserts_item_into_sorted_order(t *testing.T) {
	tree := NewTree(2)
	tree.Put(3)
	tree.Put(1)
	tree.Put(2)

	if tree.root.items[0].Val != 1 {
		t.Errorf("Expected root.items[0].Val to be 1, got %d", tree.root.items[0].Val)
	}
	if tree.root.items[1].Val != 2 {
		t.Errorf("Expected root.items[1].Val to be 2, got %d", tree.root.items[1].Val)
	}
	if tree.root.items[2].Val != 3 {
		t.Errorf("Expected root.items[1].Val to be 3, got %d", tree.root.items[2].Val)
	}
}

func Test_tree_inserts_item_into_correct_node(t *testing.T) {
	tree := NewTree(2)
	tree.Put(3)
	tree.Put(1)
	tree.Put(2)
	tree.Put(4)
	tree.Put(5)
	tree.Put(6)
	tree.Put(7)
	tree.Put(8)

	if len(tree.root.children) != 2 {
		t.Errorf("Expected root.children to have 2 children, got %d", len(tree.root.children))
	}
	if len(tree.root.children[0].items) != 2 {
		t.Errorf("Expected root.children[0].items to have 2 items, got %d", len(tree.root.children[0].items))
	}
	if len(tree.root.children[1].items) != 3 {
		t.Errorf("Expected root.children[1].items to have 3 items, got %d", len(tree.root.children[1].items))
	}

}
