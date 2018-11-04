package myintf

import (
	"myoop/myintf/a"
	"myoop/myintf/b"
)

/*
如果回圈引用就会报以下错误：
import cycle not allowed
package main
	imports iotestgo/myoop/myintf
	imports iotestgo/myoop/myintf/a
	imports iotestgo/myoop/myintf/b
	imports iotestgo/myoop/myintf/a
*/
func TestAAABBBCCC() {
	obj := a.AAA{"AAA"}
	obj1 := b.BBB{"BBB"}
	obj.DisAAA(&obj1)
	obj1.DisBBB(&obj)
}
