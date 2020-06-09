package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depthRight := maxDepth(root.Right)
	depthLeft := maxDepth(root.Left)

	if depthLeft > depthRight {
		return depthLeft + 1
	}

	return depthRight + 1
}

func main() {
	println(maxDepth(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
}
