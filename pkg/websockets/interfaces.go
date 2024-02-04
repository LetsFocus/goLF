package websockets

import "context"

type methods interface {
	Publish(ctx *context.Context)
	Subscribe(ctx *context.Context)
}
