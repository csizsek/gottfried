package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"github.com/csizsek/gottfried/common"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

type S3Operation struct {}

func (o *S3Operation) S3List(args *common.S3ListArgs, reply *common.S3ListResult) error {
	log.Printf("S3List bucket: %s", args.Bucket)
	*reply = common.S3ListResult{}
	aws_access_key := os.Getenv("AWS_ACCESS_KEY_ID")
	if aws_access_key == "" {
		log.Fatal("Environment variable AWS_ACCESS_KEY_ID is not set")
	}
	aws_secret_key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if aws_secret_key == "" {
		log.Fatal("Environment variable AWS_SECRET_ACCESS_KEY is not set")
	}
	auth := aws.Auth{AccessKey: aws_access_key, SecretKey: aws_secret_key}
	region := aws.USEast
	s3 := s3.New(auth, region)
	bucket := s3.Bucket(args.Bucket)
	list, err := bucket.List("", ",", "", 5)
	if err != nil {
		return errors.New("Unable to list bucket\n" + err.Error())
	}
	files := make([]string, len(list.Contents))
	for i, k := range list.Contents {
		files[i] = k.Key
	}
	reply.List = files
	return nil
}

func main() {
	s3Op := new(S3Operation)
	rpc.Register(s3Op)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal("Unable to listen\n" + err.Error())
	}
	http.Serve(l, nil)
}
