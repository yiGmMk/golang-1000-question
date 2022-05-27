package typetest

import (
	"log"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/yiGmMk/leetcode/golang/util"
)

/*
type NodeType1 util.TreeNode

type NodeType2 = util.TreeNode
*/
/*
 1. type T1   T2, T1和T2结构一样,但是绑定的method set 对象不同,T1不继承T2的method set,定义同一名称方法不会产生冲突
    type T1 = T2, T1是T2的是类型别名,方法用同一个绑定的method set 对象,T1和T2不能定义同一名称方法,实际上是同一类型

 1)类型别名(type aliases):
   Go now supports type aliases to support gradual code repair while moving a type between packages
   This declaration introduces an alias name T1—an alternate spelling—for the type denoted by T2; that is, both T1 and T2 denote the same type.
*/
func TestType(t *testing.T) {
	node := &util.TreeNode{Val: 1}
	var iVal interface{}
	iVal = node
	convey.Convey("type1", t, func() {
		val, ok := iVal.(*NodeType1)
		log.Println(val, ok)
		convey.So(ok, convey.ShouldBeFalse)
	})

	//
	convey.Convey("type2", t, func() {
		val, ok := iVal.(*NodeType2)
		log.Println(val, ok)
		convey.So(ok, convey.ShouldBeTrue)
	})
}

// 2. NodeType1可以
//    NodeType2不能,invalid receiver type *util.TreeNode (type not defined in this package)

/* 3.references
- https://golang.google.cn/doc/go1.9
- https://golang.google.cn/design/18130-type-alias
- https://golang.google.cn/talks/2016/refactor.article
- http://c.biancheng.net/view/25.html
- https://golang.google.cn/ref/spec
*/
