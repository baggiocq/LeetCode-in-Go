package Problem0100

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q== nil {
		return true 
	}

	if p == nil || q == nil {
		return false 
	}

	return isSameTree(p.Left, q.Left) && p.Val == q.Val && isSameTree(p.Right, q.Right)
}
