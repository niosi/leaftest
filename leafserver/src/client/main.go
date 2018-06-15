package main

import (
	"encoding/binary"
	"net"
	"io"
	"fmt"
	"log"
	"encoding/json"
	"strings"
	"bytes"
)

type RespInfo struct {
	LeafTest *LeafTest
}

type LeafTest struct {
	Name string
	Sex int
	Age int
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
		"LeafTest": {
			"Name": "leaf",
			"Sex":1,
			"Age":18
		}
	}`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)
	// 发送消息
	conn.Write(m)

	buf := make([]byte,1024) //定义一个切片的长度是1024。

	conn.Read(buf) //接收到的内容大小。

	if err != nil && err != io.EOF {  //io.EOF在网络编程中表示对端把链接关闭了。
		log.Fatal(err)
	}
	subStr := strings.Split(string(buf),".")[1]
	resp := &RespInfo{}
	json.Unmarshal(bytes.TrimRight([]byte(subStr),"\x00"),&resp)
	fmt.Printf("%#v",resp) //将接受的内容都读取出来。
	conn.Close()  //断开TCP链接。
}