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

/* 3. what are type aliases designed for ?
类型别名设计初衷

*/

/* 4. summary of a post from go talk about Codebase refactoring
talk_url: https://golang.google.cn/talks/2016/refactor.article#TOC_5.1.
- go致力于简化构建大规模的软件,这包括更大的系统,更大的代码量(google单库超十亿行代码,godoc.org为数十万个package提供服务)
- go为此限制了package的import,只导入需使用的
- 代码重构的原因:
	1. 拆分package,方便用户管理,使用
	2. 优化命名
	3. 优化依赖,减少依赖.(如将os.EOF移入io包,这样不需要os库api的可用减少对os依赖)
	4. 修改依赖关系图使某些包能被另一个包导入(在go 1前由于包括time等许多包import了os导致在os在中无法使用time包,go 1后修改了)
*/

/* 5.references
- https://golang.google.cn/doc/go1.9
- https://golang.google.cn/design/18130-type-alias
- https://golang.google.cn/talks/2016/refactor.article
- http://c.biancheng.net/view/25.html
- https://golang.google.cn/ref/spec
*/
