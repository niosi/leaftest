package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"reflect"
	"server/msg"
	"time"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	// 向当前模块（login 模块）注册消息的消息处理函数 handleTest
	handleMsg(&msg.Login{}, handleTest)
}

func handleTest(args []interface{}) {
	fmt.Printf("request start time:%s\n", time.Now().Format("2006-01-02 15:04:05"))
	// 收到的 Test 消息
	m := args[0].(*msg.Login)
	// 消息的发送者
	a := args[1].(gate.Agent)
	// 输出收到的消息的内容
	fmt.Printf("User:%s,Pwd:%s\n", m.LoginName, m.LoginPwd)
	if m.LoginName == "testUser" && m.LoginPwd == "testPwd" {
		// 给发送者回应一个消息
		a.WriteMsg(&msg.Resp{Code: 0, Msg: "succ"})
	} else {
		a.WriteMsg(&msg.Resp{Code: 10000, Msg: "fail"})
	}
	fmt.Printf("request end time:%s\n", time.Now().Format("2006-01-02 15:04:05"))
}
