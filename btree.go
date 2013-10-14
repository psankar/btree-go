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
	for i, val := range active.elements {
		fmt.Printf("Comparing %d against %d\n", x, val)
		if val == x {
			fmt.Printf("[%d] Already existing in the btree\n", x)
			return btree
		} else if val > x {
			if len(active.children) > i {
				active = active.children[i]
				goto a
			} else {
				break
			}
		}
	}

	/* Check if there is a subtree on the right of the last element */
	if len(active.children) == len(active.elements)+1 {
		active = active.children[len(active.elements)]
		goto a
	}

	/* There is enough space to insert the element in this node
	 * itself without having to create a new node */
	active.elements = insertWithinNode(active.elements, x)

b:
	if len(active.elements) > 2*btree.order {

		rightNode := &bTreeNode{}
		rightNode.elements = make([]int, btree.order)
		rightNode.children = make([]*bTreeNode, 0, 2*btree.order+1)
		copy(rightNode.elements, active.elements[btree.order+1:])

		if active.parent == nil {
			fmt.Println("Overflow in the btree root node. Should split the node")
			fmt.Println(btree.root)
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

			printbTreeNodes(newRootNode)
			return btree
		} else {
			var pos, val int

			fmt.Println("Overflow in non-root node of the b-tree")

			parent := active.parent
			/* Get the mid element */
			x = active.elements[btree.order]

			/* Remove everything after the mid-element */
			active.elements = active.elements[:btree.order]

			/* Find the suitable position for the mid element in the
			 * parent node */
			for pos, val = range parent.elements {
				if val > x {
					break
				}
			}

			if parent.elements[pos] < x {
				pos = pos + 1
				parent.elements = append(parent.elements, x)
				parent.children = append(parent.children, rightNode)
			} else {
				parent.elements = append(parent.elements, x)
				copy(parent.elements[pos+1:], parent.elements[pos:])
				parent.elements[pos] = x

				if (pos + 1) < len(parent.children) {
					pos = pos + 1
				} else {
					/* Can this condition ever occur ? */
					panic("blah")
					pos = len(parent.children)
				}
				parent.children = append(parent.children, rightNode)
				copy(parent.children[pos+1:], parent.children[pos:])
				parent.children[pos] = rightNode
			}
			rightNode.parent = parent
			active = parent
			printbTreeNodes(btree.root)
			goto b
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

	fmt.Println("--------------")

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

	//a := []int{8, 3, 10, 1, 6, 9, 4, 7, 13, 18, 15}
	a := []int{6, 1, 3, 10, 4, 7, 8, 9, 18, 12, 13, 19, 15, 22, 33, 35, 44, 70, 37, 38, 39, 50, 60, 55, 80, 90, 101, 102, 100, 110, 120}

	for _, i := range a {
		Insert(btree, i)
		fmt.Println()
	}

	fmt.Println("Elements in the tree:")
	PrintbTree(btree)

	/*

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


	*/
}
