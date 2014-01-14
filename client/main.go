package main

import (
  "flag"
	"github.com/aliafshar/runnerd"
)

func main() {
  flag.Parse()
  c, _ := runnerd.NewClient(1234)
  commandArgs := flag.Args()
  c.Request(commandArgs[0], commandArgs[1:]...)
}
