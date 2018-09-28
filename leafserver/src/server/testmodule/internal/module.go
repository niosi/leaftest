/*
package: package internal
author: Wuhaiqiang
date: 2018-06-15 上午 10:03
*/

package internal

import (
	"fmt"
	"github.com/name5566/leaf/module"
	"server/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	//log.Debug("1")
	//// 定义变量 res 接收结果
	//var res string
	//
	//skeleton.Go(func() {
	//	// 这里使用 Sleep 来模拟一个很慢的操作
	//	time.Sleep(1 * time.Second)
	//
	//	// 假定得到结果
	//	res = "3"
	//}, func() {
	//	log.Debug(res)
	//})
	//
	//log.Debug("2")
}

func (m *Module) OnDestroy() {
	fmt.Println("OnDestory")
}
