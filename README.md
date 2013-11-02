#Gottfried

##What is Gottfried

This application is not useful at all, it is just a demonstration of Go and some of its libraries, a proof of concept web service featuring a HTTP REST API, RPC, a worker with AWS operations and stuff.

##How to use

###To compile the worker
    $ export GOPATH=`pwd`
    $ go get github.com/csizsek/gottfried/worker

###To compile the service
    $ export GOPATH=`pwd`
    $ go get github.com/csizsek/gottfried/service

###To run the worker
    $ export AWS_ACCESS_KEY_ID=<your AWS access key>
    $ export AWS_SECRET_ACCESS_KEY=<your AWS secret key>
    $ ./bin/worker
It will try to bind port 8001 for the RPC server.

###To run the service
    $ ./bin/service
It will try to bind port 8000 for the HTTP service.

###To try the service
    $ curl localhost:8000/gottfried/api/v1/s3/list/<your bucket>
It will return the first 5 entries in the specified S3 bucket in JSON format.