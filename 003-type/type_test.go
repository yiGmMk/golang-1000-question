package typetest

import (
	"log"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/yiGmMk/leetcode/golang/util"
)

func TestType(t *testing.T) {
	node := util.TreeNode{Val: 1}
	convey.Convey("type1", t, func() {
		out, err := Trans1(node)
		convey.So(out, convey.ShouldNotBeNil)
		log.Println("type1 error:", err)
	})

	//
	convey.Convey("type2", t, func() {
		out, err := Trans2(node)
		convey.So(err, convey.ShouldBeNil)
		convey.So(out.Val, convey.ShouldEqual, node.Val)
		log.Println("type2 error:", err)
	})
}
