package main

import (
	"github.com/victor-cabrera/fe-calc/server/core"
)

func main() {
	p := core.NewPlatform()

	p.Success("Fire Emblem Calculator daemon running...\n")

	p.Start()
}
