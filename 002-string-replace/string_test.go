package goquestion

import (
	"fmt"
	"strings"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	origin := "   特斯拉Model X    "
	replace := strings.ReplaceAll(origin, " ", "")

	fmt.Printf("origin:%s\nreplace:%s\n", origin, replace)

	convey.Convey("string replace", t, func() {
		replaceV1 := Trim(origin, []rune{' ', '-'})
		fmt.Printf("origin:%s\nreplace_v1:%s\n", origin, replaceV1)
		convey.So(replaceV1, convey.ShouldEqual, "特斯拉ModelX")
	})
}
