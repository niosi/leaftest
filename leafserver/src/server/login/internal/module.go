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
	fmt.Println("OnInit")
}

func (m *Module) OnDestroy() {
	fmt.Println("OnDestory")
}
