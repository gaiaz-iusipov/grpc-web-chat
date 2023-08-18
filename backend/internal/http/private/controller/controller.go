package httpprivatecontroller

import (
	"sync/atomic"

	httpprivate "github.com/gaiaz-iusipov/grpc-web-chat/internal/http/private"
)

type Controller struct {
	ready atomic.Bool
}

var _ httpprivate.Controller = (*Controller)(nil)

func New() *Controller {
	return &Controller{}
}

func (c *Controller) SetReady(ready bool) {
	c.ready.Store(ready)
}
