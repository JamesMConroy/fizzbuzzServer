# FizzBuzz Server

This is a simple project to implement fizzbuzz as a server

Build the project with `go build` and you run `fizzbuzzServer`.
The server will listen on port 5000, so you can `curl localhost:5000/fizzbuzz/<number>` to get a response.

A dockerfile is also provided to run the generated binary.

Ex:

``` sh
docker run --rm -it -p 5000:5000 fizzbuzz:server
```
