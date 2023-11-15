package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	// "github.com/lonng/nano"
	// "github.com/lonng/nano/component"
	// "github.com/lonng/nano/examples/demo/tadpole/logic"
	// "github.com/lonng/nano/serialize/json"
	// "github.com/lonng/nano"
	// "github.com/lonng/nano/component"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/groups"

	"pitaya_tadpole/logic"
)

func main() {
	port := flag.Int("port", 23456, "the port to listen")
	svType := flag.String("type", "game", "the server type")
	flag.Parse()

	builder := pitaya.NewDefaultBuilder(true, *svType, pitaya.Standalone /*pitaya.Cluster*/, map[string]string{}, *config.NewDefaultBuilderConfig())
	tcp := acceptor.NewWSAcceptor(fmt.Sprintf(":%d", *port))
	builder.AddAcceptor(tcp)

	builder.Groups = groups.NewMemoryGroupService(*config.NewDefaultMemoryGroupConfig())

	app := builder.Build()
	defer app.Shutdown()

	world := logic.NewWorld(app)
	app.Register(world,
		component.WithName("world"),
		component.WithNameFunc(strings.ToLower),
	)
	manager := logic.NewManager(app)
	app.Register(manager,
		component.WithName("manager"),
		component.WithNameFunc(strings.ToLower),
	)

	// http://127.0.0.1:23456/static/
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../tadpole-client"))))
	go http.ListenAndServe(":9000", nil)

	app.Start()
}
