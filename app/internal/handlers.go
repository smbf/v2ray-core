package internal

import (
	"github.com/v2ray/v2ray-core/app"
	"github.com/v2ray/v2ray-core/proxy"
)

type InboundHandlerManagerWithContext interface {
	GetHandler(context app.Context, tag string) (proxy.InboundHandler, int)
}

type inboundHandlerManagerWithContext struct {
	context app.Context
	manager InboundHandlerManagerWithContext
}

func (this *inboundHandlerManagerWithContext) GetHandler(tag string) (proxy.InboundHandler, int) {
	return this.manager.GetHandler(this.context, tag)
}