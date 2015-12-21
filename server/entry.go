package server

import (
	"io"
)

// Server exports the common server interfaces
type Server interface {
	Start()
	Stop()
	Shutdown()
	Init()
	Output(io.Writer)
}
