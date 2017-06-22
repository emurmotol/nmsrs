package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/emurmotol/nmsrs/env"
	"github.com/emurmotol/nmsrs/model"
	"github.com/emurmotol/nmsrs/router"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	reset, _ := env.Conf.Bool("pkg.mgo.reset")
	model.Load(reset)

	// gob.Register(helper.Alert{}) // info: register the struct so encoding/gob knows about it

	port, _ := env.Conf.Int("server.port")
	addr := fmt.Sprintf(":%d", port)
	r := router.Handler()
	log.Printf("main: serving at %s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		panic(err)
	}
}
