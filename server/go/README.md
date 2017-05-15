# Installation

Install golang
Set GOPATH:
    cd to directory of this file
    export GOPATH=$(pwd)

go get github.com/gorilla/mux
go get golang.org/x/net/context
go get golang.org/x/oauth2
go get golang.org/x/oauth2/facebook
go get github.com/asaskevich/govalidator

# Run
### Make sure to run in path with view folder
cd go
go build main && ./main