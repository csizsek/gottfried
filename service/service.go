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
	dBStore            gorest.EndPoint `method:"POST" path:"/db/store" postdata:"string"`
}

func main() {
	log.SetPrefix("Gottfried Service ")
	configFile := flag.String("conf", "service.cfg", "The name of the service configuration file")
	flag.Parse()
	err := gcfg.ReadFileInto(&CFG, *configFile)
	if err != nil {
		log.Fatal("Unable to load configuration file: " + *configFile + "\n" + err.Error())
	}
	service := new(GottfriedService)
	gorest.RegisterService(service)
	http.Handle("/", gorest.Handle())
	log.Println("Service started. Listening for HTTP calls on port " + CFG.HTTP.Port)
	log.Println("Press CTRL+C to stop the service process")
	http.ListenAndServe(CFG.HTTP.Port, nil)
}
