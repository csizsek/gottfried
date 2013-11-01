package main

import (
    "net/http"
    "code.google.com/p/gorest"
)

type GottfriedService struct {
    gorest.RestService `root:"/gottfried/api/v1/"`
    s3List gorest.EndPoint `method:"GET" path:"/s3/list/{bucket:string}" output:"string"`
}

func main() {
    gorest.RegisterService(new(GottfriedService))
    http.Handle("/",gorest.Handle())    
    http.ListenAndServe(":8000",nil)
}
