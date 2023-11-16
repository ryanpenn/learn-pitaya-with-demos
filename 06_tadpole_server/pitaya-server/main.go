package main

import (
	"flag"
	"fmt"
	"net/http"
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
	flag.Parse()

	builder := pitaya.NewDefaultBuilder(true, "", pitaya.Standalone /*pitaya.Cluster*/, map[string]string{}, *config.NewDefaultBuilderConfig())
	tcp := acceptor.NewWSAcceptor(fmt.Sprintf(":%d", *port))
	builder.AddAcceptor(tcp)

	builder.Groups = groups.NewMemoryGroupService(*config.NewDefaultMemoryGroupConfig())

	app := builder.Build()
	defer app.Shutdown()

	world := logic.NewWorld(app)
	app.Register(world,
		component.WithName("World"),
	)
	manager := logic.NewManager(app)
	app.Register(manager,
		component.WithName("Manager"),
	)

	// http://127.0.0.1:9000/static/
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../tadpole-client"))))
	go http.ListenAndServe(":9000", nil)

	app.Start()
}
