package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/doublesouth/gin-scaffold/conf"
	"github.com/doublesouth/gin-scaffold/starter"
)

func main() {
	defer func() {
		glog.Flush()
	}()

	flag.Parse()
	err := conf.InitConfig()
	if err != nil {
		glog.Fatal(err)
	}

	starter.Worker()
	starter.Gin()
}
