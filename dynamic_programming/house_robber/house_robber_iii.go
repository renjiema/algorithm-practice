package house_robber

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 题3：https://leetcode-cn.com/problems/house-robber-iii
var nodeCache = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if res, ok := nodeCache[root]; ok {
		return res
	}
	// 不抢，去下一层节点
	notDo := rob(root.Left) + rob(root.Right)
	// 抢,去下下层节点
	do := root.Val
	if root.Left != nil {
		do += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		do += rob(root.Right.Left) + rob(root.Right.Right)
	}
	nodeCache[root] = max(notDo, do)
	return nodeCache[root]
}

func rob_3_2(root *TreeNode) int {
	return max(dp(root))
}
func dp(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	noRobLeft, robLeft := dp(root.Left)
	noRobRight, robRight := dp(root.Right)
	rob := root.Val + noRobLeft + noRobRight
	noRob := max(noRobLeft, robLeft) + max(noRobRight, robRight)
	return noRob, rob
}
