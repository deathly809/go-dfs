//
//	dfs is a fault tolerant distributed file system.  It allows for auto allocation of resources and
//  load balancing.
//
//

package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/deathly809/godfs/server"
)

func main() {

	isController := flag.Bool("daemon", false, "starts up as a controller")
	flag.Parse()

	var server server.Server

	if *isController {
		server = server.NewController()
	} else {
		server = server.NewWorker()
	}

	server.Init()
	go server.Start()

	// Capture SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			server.Stop()
		}
	}()

	server.Shutdown()

}
