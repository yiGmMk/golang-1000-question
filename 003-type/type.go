package typetest

import (
	"github.com/yiGmMk/leetcode/golang/util"
)

//1. we have two customize type, what is the difference?
//两种类型有何区别
type NodeType1 util.TreeNode

type NodeType2 = util.TreeNode

// node := &util.TreeNode{Val: 1}
// var iVal interface{}
// iVal = node
// val, ok := iVal.(*NodeType1)  //val,ok is ?
// val, ok := iVal.(*NodeType2)  //val,ok is ?

// 2. can we define function as below?
// 能定义下面这样的函数吗?
// func (n *NodeType1) AddLeft(l *NodeType1) {
// 	n.Left = (*util.TreeNode)(l)
// }

// func (n *NodeType2) AddLeft(l *NodeType1) {
// 	n.Left = (*util.TreeNode)(l)
// }
