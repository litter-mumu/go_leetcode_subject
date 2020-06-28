package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。


示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/

//方法1 队列
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode //定义一个队列
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		subRes := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[i]
			subRes = append(subRes, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, subRes)
		queue = queue[size:] //把上一层的节点弹出，pop
	}

	return res
}

//方法2 递归
func levelOrder1(root *TreeNode) [][]int {
	var result [][]int
	result = make([][]int, 0)
	if root == nil {
		return result
	}
	dfsHelper(root, 0, &result)
	return result
}

func dfsHelper(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}
	if len(*result) < level+1 {
		sub := make([]int, 0)
		*result = append(*result, sub)
	}
	(*result)[level] = append((*result)[level], node.Val)
	dfsHelper(node.Left, level+1, result)
	dfsHelper(node.Right, level+1, result)
}

/**
二叉树的最大深度
*/
func maxDepth(root *TreeNode) int {
	var answer int = 0
	if root == nil {
		return answer
	}
	helper(root, &answer, 0)
	return answer
}

func helper(node *TreeNode, answer *int, depth int) {
	if node == nil {
		if depth > *answer {
			*answer = depth
		}
		return
	}
	helper(node.Left, answer, depth+1)
	helper(node.Right, answer, depth+1)
}

/**
路径总和
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return helper2(root, &sum, 0)
}

func helper2(node *TreeNode, sum *int, currentSum int) bool {
	if node == nil {
		return false
	}
	if node.Val+currentSum == *sum && node.Left == nil && node.Right == nil {
		return true
	}
	return helper2(node.Left, sum, currentSum+node.Val) || helper2(node.Right, sum, currentSum+node.Val)
}

/**
对称二叉树
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3

进阶：
你可以运用递归和迭代两种方法解决这个问题吗？
*/
func isSymmetric(root *TreeNode) bool {
	return helper3(root, root)
}
func helper3(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1 == nil || tree2 == nil {
		return false
	}
	return tree1.Val == tree2.Val &&
		helper3(tree1.Left, tree2.Right) &&
		helper3(tree1.Right, tree2.Left)
}

/**
106. 从中序与后序遍历序列构造二叉树
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	//声明根
	var root *TreeNode
	//判断后序遍历的最后结点
	if len(postorder) == 0 {
		return root
	}
	//构造根
	root = new(TreeNode)
	root.Val = postorder[len(postorder)-1]
	//获取在中序遍历的位置
	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	//构造左子树和右子树
	root.Left = buildTree(inorder[:i], postorder[:i])//i=1  [9]  [9]
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	//返回
	return root

}

func main() {
	var tree = &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15, nil, nil,
			},
			&TreeNode{
				7, nil, nil,
			},
		},
	}
	fmt.Println(isSymmetric(tree))
}
