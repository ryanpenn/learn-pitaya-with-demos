package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/groups"

	"pitaya_tadpole/logic"
)

func configureBackend(app pitaya.Pitaya) {
	store := logic.NewStore()
	// RegisterRemote
	app.RegisterRemote(store,
		component.WithName("store"),
		component.WithNameFunc(strings.ToLower),
	)
}

func configureFrontend(app pitaya.Pitaya) {
	world := logic.NewWorld(app)
	app.Register(world,
		component.WithName("World"),
	)

	manager := logic.NewManager(app)
	app.Register(manager,
		component.WithName("Manager"),
	)
}

func main() {
	port := flag.Int("port", 23456, "the port to listen")
	front := flag.Bool("front", true, "is frontend server")
	flag.Parse()

	svType := ""
	if !*front {
		svType = "store"
	}

	// 以集群的模式运行
	builder := pitaya.NewDefaultBuilder(*front, svType, pitaya.Cluster, map[string]string{}, *config.NewDefaultBuilderConfig())
	if *front {
		ws := acceptor.NewWSAcceptor(fmt.Sprintf(":%d", *port))
		builder.AddAcceptor(ws)
	}

	builder.Groups = groups.NewMemoryGroupService(*config.NewDefaultMemoryGroupConfig())

	app := builder.Build()
	defer app.Shutdown()

	if *front {
		configureFrontend(app)
		// http://127.0.0.1:9000/static/
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../tadpole-client"))))
		go http.ListenAndServe(":9000", nil)
	} else {
		configureBackend(app)
	}

	app.Start()
}
