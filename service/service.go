package main

import (
	"code.google.com/p/gorest"
	"net/http"
)

type GottfriedService struct {
	gorest.RestService `root:"/gottfried/api/v1/"`
	s3List             gorest.EndPoint `method:"GET" path:"/s3/list/{bucket:string}" output:"string"`
	//s3Store gorest.EndPoint `method:"POST" path:"/s3/store/{bucket:string}/{file:string}" postdata:"string"`
}

func main() {
	gorest.RegisterService(new(GottfriedService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":8000", nil)
}
