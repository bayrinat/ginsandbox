package main

import (
	"flag"

	"github.com/bayrinat/geeksteam/http"
)

const (
	defaultAddr = ":8080"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "address", defaultAddr, "bind address")
}

func main() {
	flag.Parse()

	router := http.NewRouter()

	router.Initialize()
	router.Run(addr)
}
