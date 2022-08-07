# B+Tree

a simple b+tree implementation in Go for learning purpose

```mermaid
graph TD
10[#10: 16, 32, ] --> 2 
2[#2: 4, 8, 12, ] --> 0 
0[#0: 1, 2, 3, ] 
2[#2: 4, 8, 12, ] --> 3 
3[#3: 5, 6, 7, ] 
2[#2: 4, 8, 12, ] --> 4 
4[#4: 9, 10, 11, ] 
2[#2: 4, 8, 12, ] --> 5 
5[#5: 13, 14, 15, ] 
10[#10: 16, 32, ] --> 11 
11[#11: 20, 24, 28, ] --> 6 
6[#6: 17, 18, 19, ] 
11[#11: 20, 24, 28, ] --> 7 
7[#7: 21, 22, 23, ] 
11[#11: 20, 24, 28, ] --> 8 
8[#8: 25, 26, 27, ] 
11[#11: 20, 24, 28, ] --> 9 
9[#9: 29, 30, 31, ] 
10[#10: 16, 32, ] --> 16 
16[#16: 36, 40, 44, ] --> 12 
12[#12: 33, 34, 35, ] 
16[#16: 36, 40, 44, ] --> 13 
13[#13: 37, 38, 39, ] 
16[#16: 36, 40, 44, ] --> 14 
14[#14: 41, 42, 43, ] 
16[#16: 36, 40, 44, ] --> 15 
15[#15: 45, 46, 47, 48, 49, 50, ] 

```

## Todos
- [x] Insertion
- [x] Find
- [ ] Deletion
- [ ] Scan
- [ ] Size
## Usage
```go
tree := NewTree(3)

for i := 1; i <= 10; i++ {
  tree.Put(i)
}

fmt.Println(tree.Print())
// Node 2: 4, 
//  Node 0: 1, 2, 3, 
//  Node 3: 5, 6, 7, 8, 9, 10, 

tree.Find(9) // true
tree.Find(12) // false
```

## Debug
You can use `tree.Mermaid()` to print the mermaid syntax list of the tree. also you can use **Mermaid** `Create()` and `RenderAndDisplay()` to get a valid md file with valid mermaid syntax

```go
m := Mermaid{}
m.Create(tree.Mermaid())
m.RenderAndDisplay()
```