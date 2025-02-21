package main

import (
	"fmt"

	"github.com/nirmata/license-guardian/pkg/guard"
)

func main() {
	fmt.Println("Hello world")
	g := guard.New()
	g.Orchestrate()
}
