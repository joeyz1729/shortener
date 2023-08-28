package connect

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	convey.Convey("base example", t, func() {
		url := "https://www.baidu.com/?tn=80035161_1_dg"
		get := Get(url)
		convey.So(get, convey.ShouldEqual, true)
		convey.ShouldBeTrue(get)
	})

	convey.Convey("invalid example", t, func() {
		url := "./ffsasfw"
		get := Get(url)
		// convey.So(get, convey.ShouldEqual, true)
		convey.ShouldBeFalse(get)
	})
}
