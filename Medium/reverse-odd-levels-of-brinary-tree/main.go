package main

import "fmt"

// 2415. Reverse Odd Levels of Binary Tree
// Solved
// Medium
// Topics
// Companies
// Hint
// Given the root of a perfect binary tree, reverse the node values at each odd level of the tree.

// For example, suppose the node values at level 3 are [2,1,3,4,7,11,29,18], then it should become [18,29,11,7,4,3,1,2].
// Return the root of the reversed tree.

// A binary tree is perfect if all parent nodes have two children and all leaves are on the same level.

// The level of a node is the number of edges along the path between it and the root node.

// Example 1:
// 1                                               1
// / \                                            / \
// 2   3                                         3  2
// / \ / \	                                     / \ / \
// 4  5 6  7                                   4  5 6  7

// Input: root = [1,2,3,4,5,6,7]
// Output: [1,3,2,4,5,6,7]

// Input: root = [7,13,11]
// Output: [7,11,13]
// Explanation:
// The nodes at level 1 are 13, 11, which are reversed and become 11, 13.
// Example 3:

// Input: root = [0,1,2,0,0,0,0,1,1,1,1,2,2,2,2]
// Output: [0,2,1,0,0,0,0,2,2,2,2,1,1,1,1]
// Explanation:
// The odd levels have non-zero values.
// The nodes at level 1 were 1, 2, and are 2, 1 after the reversal.
// The nodes at level 3 were 1, 1, 1, 1, 2, 2, 2, 2, and are 2, 2, 2, 2, 1, 1, 1, 1 after the reversal.

// Definition for a binary tree node.
 type TreeNode struct {
     Val int
      Left *TreeNode
    Right *TreeNode
  }

 func reverseOddLevels(root *TreeNode) *TreeNode {
    
    reverse(1, root.Left, root.Right)
    return root
}

func reverse( level int, left *TreeNode, right *TreeNode){
    if left== nil || right ==nil{
        return
    }
    if level%2 !=0{
        left.Val, right.Val = right.Val, left.Val
    }
    reverse(level+1, left.Left, right.Right)
    reverse(level+1, right.Left, left.Right)
}
// Example usage
func main() {
	// Constructing a sample binary tree:
	//         1
	//        / \
	//       2   3
	//      / \ / \
	//     4  5 6  7
	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7}}}

	fmt.Println("Original Tree:")
	printTree(root, 0)

	// Call the reverseOddLevels function to reverse values at odd levels.
	root = reverseOddLevels(root)

	fmt.Println("\nTree After Reversing Odd Levels:")
	printTree(root, 0)
}
// Helper function to print the tree for debugging purposes
func printTree(root *TreeNode, level int) {
	if root == nil {
		return
	}
	printTree(root.Right, level+1)
	fmt.Printf("%s%d\n", string(level*4), root.Val)
	printTree(root.Left, level+1)
}