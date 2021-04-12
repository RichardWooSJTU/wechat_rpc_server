package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"wechatpro/gen-go/wechat"
)

const (
	auth = "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxNzUyMTc3NjU4MiJ9.OGkdR3cKxGrjcWkjuhGVjxyb5dXxDkNcZ6BsoVxU1z3-NHD_MTKPkeXieE_5BVosmzeaH86ci0emJQG-rkPLXg"
)

var wId string

type WechatServer struct {

}

func (w *WechatServer)  Send(ctx context.Context, option int32, content string) (_r string, _err error) {
	return "", nil
}

func (w *WechatServer) FetchGroups(ctx context.Context) (_r []*wechat.Group, _err error) {
	err := InitAddressList(wId)
	if err != nil {
		return nil, err
	}
	groups, err := QueryGroups(wId)
	if err != nil {
		return  nil, err
	}
	err = GetGroupDetail(wId,  groups)
	if err != nil {
		return nil, err
	}
	
	return groups, nil
}

func main() {
	//确认登录，如果没有报错并返回
	fmt.Println("请输入wId")
	_, err := fmt.Scanln(&wId)
	if err != nil || wId == ""{
		panic(err)
		return
	}
	err = VerifyWechatOnline(wId)
	if err != nil {
		panic(err)
		return
	}

	transport, err := thrift.NewTServerSocket(":9876")
	if err != nil {
		panic(err)
	}
	handler := &WechatServer{}
	processor := wechat.NewWechatProcessor(handler)

	transportFactory := thrift.NewTBufferedTransportFactory(10240)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
		)
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
