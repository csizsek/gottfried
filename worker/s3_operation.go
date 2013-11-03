package main

import (
	"errors"
	"github.com/csizsek/gottfried/common"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"log"
	"os"
)

type S3Operation struct {}

func (o *S3Operation) S3List(arg *common.S3ListArg, reply *common.S3ListResult) error {
	log.Printf("S3List bucket: %s", arg.Bucket)
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
	bucket := s3.Bucket(arg.Bucket)
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
