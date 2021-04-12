package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"wechatpro/gen-go/wechat"
)

type WechatServer struct {

}

func (w *WechatServer)  Send(ctx context.Context, option int32, content string) (_r string, _err error) {
	return fmt.Sprintf("option choosed is %v and content is %v", option, content), nil
}

func (w *WechatServer) FetchGroups(ctx context.Context) (_r []*wechat.Group, _err error) {
	group := wechat.Group{
		GroupID:   "1",
		GroupName: "我是你",
	}
	_r = append(_r, &group)

	return _r, nil
}

func main() {
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
