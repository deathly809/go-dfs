//
//	dfs is a fault tolerant distributed file system.  It allows for auto allocation of resources and
//  load balancing.
//
//

package main

import (
	"github.com/deathly809/gofs/concrete"
	"flag"
	"log"
)

func main() {
	isController := flag.Bool("daemon", false, "starts up a controller")
	if *isController {
		log.Print("Controller started...")
	}else {
        log.Print("Worker started...")
    }
    
	fs,err := concrete.Open("test")
	if err != nil {
		log.Fatal(err)
	}
    defer fs.Shutdown()
}
