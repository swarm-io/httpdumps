# httpdumps

A simple Go utility to output the entirety of an http request.  

By default it will listen on port 8801 and log everything made to /dumps.  It can be run via Go or via Docker as so:

`cd src; go run server.go ; cd -`

`docker run swarm-io/httpdumps --rm httpdumps`

You can use a service like [ngrok](https://ngrok.com/) to use with some external service. 