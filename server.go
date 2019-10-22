package main

import (
	"github.com/victor-cabrera/fe-calc/server/core"
)

func main() {
	p := core.NewPlatform()

	p.Success("Server running...\n")

	p.Start()
}
