package helpers

import (
	"github.com/astaxie/beego"
	"strings"
)

func RegisterTemplateHelpers() {
	// nl2br
	beego.AddFuncMap("nl2br", func(in string) string {
		return strings.Replace(in, "\n", "<br>", -1)
	})

	// 奇数判定
	beego.AddFuncMap("isOdd", func(number int) bool {
		return number%2 != 0
	})
}
