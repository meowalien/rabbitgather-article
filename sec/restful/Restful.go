package restful

import (
	"context"
	"github.com/meowalien/rabbitgather-article/conf"
)




type RestfulHandler struct {
	Debug bool
	Config conf.RestfulHandlerConfiguration
}

func (h *RestfulHandler) Start(ctx context.Context) {
	//log.DEBUG.Println("APIServer listen on : ", h.Config.Port)

}

func (h *RestfulHandler) Stop(ctx context.Context) error {
	return nil
}