package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/emurmotol/nmsrs.v4/env"
	"github.com/emurmotol/nmsrs.v4/helper"
	"github.com/emurmotol/nmsrs.v4/model"
	"github.com/emurmotol/nmsrs.v4/router"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	reset, _ := env.Conf.Bool("pkg.gorm.reset")
	model.Load(reset)

	gob.Register(helper.Alert{}) // info: register the struct so encoding/gob knows about it

	port, _ := env.Conf.Int("server.port")
	addr := fmt.Sprintf(":%d", port)
	r := router.Handler()
	log.Printf("main: serving at %s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		panic(err)
	}
}
