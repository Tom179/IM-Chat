package main

import (
	"flag"
	"fmt"

	"Im-chat/Chat/apps/user/api/internal/config"
	"Im-chat/Chat/apps/user/api/internal/handler"
	"Im-chat/Chat/apps/user/api/internal/svc"
	"Im-chat/Chat/pkg/xresp"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/dev/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandlerCtx(xresp.ErrHandler(c.Name))
	httpx.SetOkHandler(xresp.OkHandler) 

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
