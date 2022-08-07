package main

import (
	"fmt"
)

func main() {
	tree := NewTree(3)

	for i := 1; i <= 50; i++ {
		fmt.Printf("Inserting %d\n", i)
		tree.Put(i)
	}

	fmt.Println(tree.Print())

	m := Mermaid{}
	m.Create(tree.Mermaid())
	m.RenderAndDisplay()

	fmt.Printf("Finding 50 in tree: %v \n", tree.Find(50))
	fmt.Printf("Finding 43 in tree: %v \n", tree.Find(43))
	fmt.Printf("Finding 77 in tree: %v \n", tree.Find(77))
	fmt.Printf("Finding 76 in tree: %v \n", tree.Find(76))
}
