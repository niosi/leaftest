package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

type RespInfo struct {
	LeafTest *LeafTest
	Resp     *RespModel
}

type LeafTest struct {
	Name string
	Sex  int
	Age  int
}

//返回消息
type RespModel struct {
	Code int
	Msg  interface{}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	//data := []byte(`{
	//	"Hello": {
	//		"Name": "leaf"
	//	}
	//}`)

	data := []byte(`{
		"Login": {
			"LoginName": "testUser",
			"LoginPwd":"testPwd"
		}
	}`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)
	// 发送消息
	conn.Write(m)

	buf := make([]byte, 1024) //定义一个切片的长度是1024。

	conn.Read(buf) //接收到的内容大小。

	if err != nil && err != io.EOF { //io.EOF在网络编程中表示对端把链接关闭了。
		log.Fatal(err)
	}

	fmt.Println(string(buf))

	resp := &RespInfo{}
	json.Unmarshal(bytes.Trim(buf, "\x00"), &resp)
	fmt.Printf("%#v", resp.Resp) //将接受的内容都读取出来。
	conn.Close()            //断开TCP链接。
}
