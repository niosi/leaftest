/*
package: package internal
author: Wuhaiqiang
date: 2018-06-15 上午 10:03
*/

package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"reflect"
	"server/msg"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.LeafTest{}, handleLeafTest)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLeafTest(args []interface{}) {
	// 收到的 LeafTest 消息
	m := args[0].(*msg.LeafTest)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %#v", m)

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.LeafTest{
		Name: "zhang",
		Sex:  1,
		Age:  18,
	})
}
