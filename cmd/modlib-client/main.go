package main

import (
	"fmt"

	"github.com/vizicist/gospout"
	"github.com/vizicist/gospout/clientlib"
)

func main() {
	fmt.Println("Running client")
	fmt.Println("Config:", gospout.Config())
	fmt.Println(clientlib.Hello())
}
