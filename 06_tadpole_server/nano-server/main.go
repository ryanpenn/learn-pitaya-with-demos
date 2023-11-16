package main

import (
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/serialize/json"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"os"

	"tadpole/logic"
)

func main() {
	app := cli.NewApp()

	app.Name = "tadpole"
	app.Author = "nano authors"
	app.Version = "0.0.1"
	app.Copyright = "nano authors reserved"
	app.Usage = "tadpole"

	// flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Value: ":23456",
			Usage: "game server address",
		},
	}

	app.Action = serve

	app.Run(os.Args)
}

func serve(ctx *cli.Context) error {
	components := &component.Components{}
	components.Register(logic.NewManager())
	components.Register(logic.NewWorld())

	// register all service
	options := []nano.Option{
		nano.WithIsWebsocket(true),
		nano.WithLabel("game"),
		nano.WithComponents(components),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithCheckOriginFunc(func(_ *http.Request) bool { return true }),
	}

	//nano.EnableDebug()
	log.SetFlags(log.LstdFlags | log.Llongfile)

	// http://127.0.0.1:23456/static/
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../tadpole-client"))))

	addr := ctx.String("addr")
	nano.Listen(addr, options...)
	return nil
}
