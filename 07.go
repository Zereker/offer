package main

/*
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

 

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
 

限制：

0 <= 节点个数 <= 5000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}

	var root TreeNode
	if len(preorder) == 1 {
		root.Val = preorder[0]
		return &root
	}

	var leftBorder, rightBorder int
	for _, v := range inorder {
		if v == root.Val {
			break
		}

		leftBorder++
	}

	rightBorder = len(preorder) - leftBorder - 1
	if leftBorder > 0 {
		root.Left = buildTree(preorder[1:leftBorder+1], inorder[0:leftBorder])
	}

	if rightBorder > 0 {
		root.Right = buildTree(preorder[leftBorder+1:], inorder[leftBorder+1:])
	}

	return &root
}
