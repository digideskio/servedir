servedir
========

A simple HTTP server to host a single static directory.

#### INSTALL
Have Go installed? Is $GOPATH/bin on your $PATH? 

Just run `go get github.com/jprobinson/servedir`  

#### USAGE

* Serve the current directory on localhost:8080
`servedir .`

* Serve the current directory on localhost:80
`servedir . 80`

* Serve some random directory on localhost:7070
`servedir ~/some/random/dir 7070`

  

