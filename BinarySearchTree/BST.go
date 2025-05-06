package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Insert(value int) {
	if value >= t.Val {
		if t.Right == nil {
			t.Right = &TreeNode{Val: value}
		} else {
			t.Right.Insert(value)
		}
	} else {
		if t.Left == nil {
			t.Left = &TreeNode{Val: value}
		} else {
			t.Left.Insert(value)
		}
	}
}

func (t *TreeNode) Find(value int) *TreeNode {
	if t == nil {
		return nil
	}
	if value == t.Val {
		return t
	} else if value > t.Val {
		return t.Right.Find(value)
	} else {
		return t.Left.Find(value)
	}
}

func (t *TreeNode) DisplayPreOrder() {
	if t != nil {
		fmt.Print(t.Val, " ")
		t.Left.DisplayPreOrder()
		t.Right.DisplayPreOrder()
	}
}

func main() {
	root := &TreeNode{Val: 10}
	values := []int{5, 15, 3, 7, 12, 18}

	for _, v := range values {
		root.Insert(v)
	}

	fmt.Println("Pre-order traversal:")
	root.DisplayPreOrder()
	fmt.Println()

	fmt.Println("\nSearching for 7:")
	found := root.Find(7)
	if found != nil {
		fmt.Println("Found:", found.Val)
	} else {
		fmt.Println("Not found")
	}
	fmt.Println("\nSearching for 100:")
	found = root.Find(100)
	if found != nil {
		fmt.Println("Found:", found.Val)
	} else {
		fmt.Println("Not found")
	}
}
