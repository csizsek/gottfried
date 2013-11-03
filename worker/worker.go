package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"code.google.com/p/gcfg"
)

type WorkerConfig struct {
	RPC struct {
		Protocol string
		Port string
	}
	DB struct {
		Host string
		Port string
		User string
		Password string
		Database string
		Table string
	}
}

var CFG WorkerConfig

func main() {
	log.SetPrefix("Gottfried Worker ")
	configFile := flag.String("conf", "worker.cfg", "The name of the worker configuration file")
	flag.Parse()
	err := gcfg.ReadFileInto(&CFG, *configFile)
	if err != nil {
		log.Fatal("Unable to load configuration file: " + *configFile + "\n" + err.Error())
	}
	s3Op := new(S3Operation)
	rpc.Register(s3Op)
	dbOp := new(DBOperation)
	rpc.Register(dbOp)
	rpc.HandleHTTP()
	l, err := net.Listen(CFG.RPC.Protocol, CFG.RPC.Port)
	if err != nil {
		log.Fatal("Unable to listen\n" + err.Error())
	}
	log.Println("Worker started. Listening for RPC calls on port " + CFG.RPC.Port)
	log.Println("Press CTRL+C to stop the worker process")
	http.Serve(l, nil)
}
