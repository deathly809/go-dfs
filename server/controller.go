package server

import (
	"io"
	"log"

	"github.com/deathly809/gofs/concrete"
)

func loadFiles(filename string) {

	fs, err := concrete.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Shutdown()
}

type controller struct {
	lg *log.Logger
}

func (c *controller) Init() {
}

func (c *controller) Start() {
	c.lg.Print("Controller started...")
	loadFiles("files")
}

func (c *controller) Stop() {
}

func (c *controller) Shutdown() {
}

func (c *controller) Output(out io.Writer) {
	c.lg = log.New(out, "Controller", log.Ldate|log.Ltime)
}

// Create instantiates and returns a new DFS Controller
func Create() Server {
	return &controller{}
}
