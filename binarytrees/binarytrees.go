package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        Walk(t.Left, ch)
    }
    ch <- t.Value
    if t.Right != nil {
        Walk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    w1, w2 := make(chan int), make(chan int)
    go Walk(t1, w1)
    go Walk(t2, w2)
    for i:=0; i<10; i++ {
        if v, v2 := <-w1, <-w2; v != v2 {
            return false
        }
    }
    return true
}

func main() {
    first := Same(tree.New(1), tree.New(1))
    second := Same(tree.New(1), tree.New(2))
    fmt.Printf("tree.New(1) == tree.New(1): %t\n", first)
    fmt.Printf("tree.New(1) == tree.New(2): %t", second)
}