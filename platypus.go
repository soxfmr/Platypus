package main

import (
	"Platypus/lib/cli/dispatcher"
	"Platypus/lib/context"
)

func main() {
	// Create context
	context.CreateContext()
	// Run main loop
	dispatcher.Run()
}
