##Gottfried

This application is not useful at all, it is just a demonstration of Go and some of its libraries, a proof of concept web service featuring a HTTP REST API, RPC, a worker with AWS operations and stuff.

##How to use

To try Gotffried you need the following things:

   * Go compiler
   * AWS S3 credentials
   * A MySql server

###To compile the worker

    $ export GOPATH=`pwd`
    $ go get github.com/csizsek/gottfried/worker

###To compile the service

    $ export GOPATH=`pwd`
    $ go get github.com/csizsek/gottfried/service

###To run the worker

    $ export AWS_ACCESS_KEY_ID=<your AWS access key>
    $ export AWS_SECRET_ACCESS_KEY=<your AWS secret key>
    $ ./bin/worker -conf <path to your config file>
    
It will try to bind the port specified in the configuration file for the RPC server.

The worker requires a running MySql server with a table that has a single varchar column.

You can find a sample configuration file called worker.cfg in the etc/ directory.

###To run the service

    $ ./bin/service -conf <path to your config file>
    
It will try to bind the port specified in your configuration file for the HTTP service.

You can find a sample configuration file called service.cfg in the etc/ directory.

###To try the service

For example:

    $ curl localhost:8000/gottfried/api/v1/s3/list/<your bucket>
    
It will return the first 5 entries in the specified S3 bucket in JSON format.

Or:

    $ curl localhost:8000/gottfried/api/v1/db/store -X POST -d "hello"
    
It will store the text "hello" in the database specified in your configuration file.