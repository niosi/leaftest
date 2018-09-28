package msg

import (
	"github.com/name5566/leaf/network/json"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = json.NewProcessor()

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&Hello{})
	Processor.Register(&LeafTest{})
	Processor.Register(&Login{})
	Processor.Register(&Resp{})
}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
type Hello struct {
	Name string
}

//测试通信
type LeafTest struct {
	Name string
	Sex  int
	Age  int
}

//LoginInfo
type Login struct {
	LoginName string
	LoginPwd  string
}

//返回消息
type Resp struct {
	Code int
	Msg interface{}
}
