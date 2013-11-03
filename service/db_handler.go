package main

import (
	"github.com/csizsek/gottfried/common"
	"log"
	"net/rpc"
)

func (service GottfriedService) DBStore(PostData string) {
	log.Printf("DBStore data: %s\n", PostData)
	client, err := rpc.DialHTTP(CFG.RPC.Protocol, CFG.RPC.Host + CFG.RPC.Port)
	if err != nil {
		log.Fatal("Unable to connect to RPC server\n" + err.Error())
	}
	arg := common.DBStoreArg{Data: PostData}
	reply := common.DBStoreResult{}
	err = client.Call("DBOperation.DBStore", arg, &reply)
	if err != nil {
		log.Fatal("Unable to call RPC method\n" + err.Error())
	}
}
