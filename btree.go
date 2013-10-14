/*
 * A thread unsafe btree that works only on integers
 *
 * Author:	Sankar <sankar.curiosity@gmail.com>
 * License:	Creative Commons Zero License
 */
package main

import (
	"fmt"
)

type bTree struct {
	root  *bTreeNode
	order int
}

type bTreeNode struct {
	elements []int
	children []*bTreeNode
	parent   *bTreeNode
}

func InitializebTree(order int) (*bTree, error) {
	if order < 1 {
		return nil, nil //error("Invalid order given for bTree")
	}

	var btree bTree
	btree.order = order

	var root bTreeNode

	/* We need space for only 2*order elements below, but we add an extra
	 * space for holding the item before splitting a node in case of overflow */
	root.elements = make([]int, 0, 2*order+1)
	root.children = make([]*bTreeNode, 0, 2*order+1)

	/* In C/C++ this will cause a dangling pointer, but perfectly valid in
	 * Go, due to the refcounting */
	btree.root = &root
	return &btree, nil
}

/* This function assumes that the slice will not overflow and has enough
 * capacity. The overflow situations should be handled by the callers */
func insertWithinNode(s []int, x int) []int {

	var i, v int

	if len(s) == 0 {
		return append(s, x)
	} else {
		for i, v = range s {
			if v > x {
				break
			}
		}

		if s[i] < x {
			return append(s, x)
		} else {
			s = append(s, x)
			copy(s[i+1:], s[i:])
			s[i] = x
			return s
		}
	}
}

func Insert(btree *bTree, x int) *bTree {
	active := btree.root

a:
	for _, i := range active.elements {
		fmt.Printf("Comparing %d against %d\n", x, i)
		if i == x {
			fmt.Printf("[%d] Already existing in the btree\n", x)
			return btree
		} else if i > x {
			if len(active.children) >= i {
				active = active.children[i]
				goto a
			} else {
				break
			}
		}
	}

	/* There is enough space to insert the element in this node
	* itself without having to create a new node */
	active.elements = insertWithinNode(active.elements, x)

	//Assert len(active.elements) <= 2*btree.order+1

	if len(active.elements) == 2*btree.order+1 {
		fmt.Println("Overflow in a btree node. Should split the node")

		rightNode := &bTreeNode{}
		rightNode.elements = make([]int, btree.order)
		rightNode.children = make([]*bTreeNode, 0, 2*btree.order+1)
		copy(rightNode.elements, active.elements[btree.order+1:])

		if active.parent == nil {
			/* root node */
			newRootNode := &bTreeNode{}
			newRootNode.elements = make([]int, 0, 2*btree.order+1)
			newRootNode.children = make([]*bTreeNode, 0, 2*btree.order+1)

			newRootNode.elements = append(newRootNode.elements, active.elements[btree.order])
			/* Remove elements after btree.order's position from
			 * active and make active the new left subtree */
			active.elements = active.elements[:btree.order]

			newRootNode.children = append(newRootNode.children, active)
			active.parent = newRootNode
			newRootNode.children = append(newRootNode.children, rightNode)
			rightNode.parent = newRootNode

			btree.root = newRootNode
			return btree
		}

	} else {
		fmt.Print("Inserted now the new array is ")
		fmt.Println(active.elements)
	}

	return btree
}

func PrintbTree(btree *bTree) {
	printbTreeNodes(btree.root)
}

func printbTreeNodes(active *bTreeNode) {

	for _, i := range active.elements {
		fmt.Println(i)
	}

	for _, child := range active.children {
		printbTreeNodes(child)
	}
}

/* Given a btree node, find the element X if it exists in the subtree(btree) */
func Find(btree *bTree, x int) *bTree {
	return btree
}

func Delete(btree *bTree, x int) *bTree {
	return btree
}

func main() {
	var btree *bTree = nil

	btree, _ = InitializebTree(3)

	a := []int{8, 3, 10, 1, 6, 9, 4, 7, 13, 18}

	for _, i := range a {
		Insert(btree, i)
		fmt.Println()
	}

	fmt.Println("Elements in the tree:")
	PrintbTree(btree)

	for _, i := range []int{1, 14, 3} {
		if Find(btree, i) != nil {
			fmt.Printf("Found %d in the Tree\n", i)
		} else {
			fmt.Printf("%d is not found in the tree\n", i)
		}
	}

	for _, i := range []int{14, 3} {
		btree = Delete(btree, i)
		fmt.Printf("Deleted %d\n", i)
	}
	PrintbTree(btree)
}
