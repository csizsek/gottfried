package main

import (
	"encoding/json"
	"github.com/csizsek/gottfried/common"
	"log"
	"net/rpc"
)

func (service GottfriedService) S3List(bucket string) string {
	log.Printf("S3List bucket: %s\n", bucket)
	client, err := rpc.DialHTTP(CFG.RPC.Protocol, CFG.RPC.Host + CFG.RPC.Port)
	if err != nil {
		log.Fatal("Unable to connect to RPC server\n" + err.Error())
	}
	args := common.S3ListArg{Bucket: bucket}
	reply := common.S3ListResult{}
	err = client.Call("S3Operation.S3List", args, &reply)
	if err != nil {
		log.Fatal("Unable to call RPC method\n" + err.Error())
	}
	list, err := json.MarshalIndent(reply.List, "", "  ")
	if err != nil {
		log.Fatal("Unable to json marshal the result\n" + err.Error())
	}
	return string(list)
}
