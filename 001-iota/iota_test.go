package goquestion

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestIota(t *testing.T) {

	fmt.Println(second)

	convey.Convey("iota", t, func() {
		convey.So(second, convey.ShouldEqual, 1)
	})
}
