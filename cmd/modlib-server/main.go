package main

import (
	"fmt"

	"github.com/vizicist/gospout"
	"github.com/vizicist/gospout/internal/auth"
	"github.com/vizicist/gospout/serverlib"
)

func main() {
	fmt.Println("Running server")
	fmt.Println("Config:", gospout.Config())
	fmt.Println("Auth:", auth.GetAuth())
	fmt.Println(serverlib.Hello())
}
