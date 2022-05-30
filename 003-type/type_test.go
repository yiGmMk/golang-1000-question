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

	// output type
	var t1, t2 = NodeType1{}, NodeType2{}
	log.Printf("\n t1_type:%T\n t2_type:%T", t1, t2)
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
- 逐步进行代码重构
	一.atomic code repair,api改动和code repair在一个大提交中完成
	二.gradual code repair,逐步重构
	三.为什么逐步重构
		1. codebase api的改动对于使用方影响很大,每一个微小的改动,逐级放大,到最终使用方可能有100倍,1000倍的改动,且在实践中
	automic code repair难以执行,大量的代码重构需要在一个提交实现,让代码实现,review变得异常困难.
		2. 通过分段执行:add new api(interchangeable with old API),code repair,remove old api,逐步重构.
			- 更大工作量,新旧api需要能同时工作.
			- 但整个流程相对简单,可以逐步完成.提交中的代码也更少,冲突的概率变小
- type类型的重构
	1. 通过gradual code repair的方式将一个type迁移到另一个package,通过type T1 T2是行不通的,即使是interface,go仍认为这是不同的类型,
在一个方法返回T2类型和T1类型被视为不同,这就导致了无法实现interface的方法(类型不同,函数前面不同).
	这里引出了问题,就是我们如何实现让一个类型和另一个interchangeable(相互可替换,等同)?atomic code repair,或者强制要求使用方在后一版本修复代码?
	为了解决这个问题(将类型从一个package逐步迁移到里一个package),go在1.9引入了类型别名(type alias).在go 1.9的release说明中有
	Go now supports type aliases to support gradual code repair while moving a type between packages.
*/

/* 5.references
- https://golang.google.cn/doc/go1.9
- https://golang.google.cn/design/18130-type-alias
- https://golang.google.cn/talks/2016/refactor.article
- http://c.biancheng.net/view/25.html
- https://golang.google.cn/ref/spec
*/
