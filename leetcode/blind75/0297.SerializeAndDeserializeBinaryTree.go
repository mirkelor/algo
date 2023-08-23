package blind75

import (
	"strconv"
	"strings"

	"github.com/mirkelor/algo/common"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *common.TreeNode) string {
	res := make([]string, 0)
	var dfs func(node *common.TreeNode)
	dfs = func(node *common.TreeNode) {
		if node == nil {
			res = append(res, "N")
			return
		}
		res = append(res, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *common.TreeNode {
	values := strings.Split(data, ",")
	i := 0
	var dfs func() *common.TreeNode
	dfs = func() *common.TreeNode {
		if values[i] == "N" {
			i++
			return nil
		}
		val, _ := strconv.Atoi(values[i])
		i++
		node := &common.TreeNode{
			Val:   val,
			Left:  dfs(),
			Right: dfs()}
		return node
	}
	return dfs()
}
