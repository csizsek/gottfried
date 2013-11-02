package main

import (
	"code.google.com/p/gcfg"
	"code.google.com/p/gorest"
	"flag"
	"log"
	"net/http"
)

type ServiceConfig struct {
	RPC struct {
		Protocol string
		Host string
		Port string
	}
	HTTP struct {
		Port string
	}
}

var CFG ServiceConfig

type GottfriedService struct {
	gorest.RestService `root:"/gottfried/api/v1/"`
	s3List             gorest.EndPoint `method:"GET" path:"/s3/list/{bucket:string}" output:"string"`
	//s3Store gorest.EndPoint `method:"POST" path:"/s3/store/{bucket:string}/{file:string}" postdata:"string"`
}

func main() {
	configFile := flag.String("conf", "service.cfg", "The name of the service configuration file")
	flag.Parse()
	err := gcfg.ReadFileInto(&CFG, *configFile)
	if err != nil {
		log.Fatal("Unable to load configuration file: " + *configFile + "\n" + err.Error())
	}
	service := new(GottfriedService)
	gorest.RegisterService(service)
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(CFG.HTTP.Port, nil)
}
