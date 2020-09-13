# httpdumps

A simple Go utility to output the entirety of an http request.  

By default it will listen on port 8801 and log everything.  It can be run via Go or via Docker as so:

`cd src; go run server.go ; cd -`

`docker run swarm-io/httpdumps --rm httpdumps`

