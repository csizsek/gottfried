package main

import (
	"encoding/json"
	"log"
	"net/rpc"	
	"github.com/csizsek/gottfried/common"
)

func(serv GottfriedService) S3List(bucket string) string {
	log.Printf("S3List bucket: %s\n", bucket)
    client, err := rpc.DialHTTP("tcp", "localhost:8001")
    if err != nil {
    	log.Fatal("Unable to connect to RPC server\n" + err.Error())
    }
    args := common.S3ListArgs{Bucket: bucket}
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

// func(serv GottfriedService) S3Store(postdata, bucket, file string) {
// 	log.Printf("S3Store bucket: %s file: %s\n", bucket, file)
// }
