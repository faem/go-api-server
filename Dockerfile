#------------------------Using busybox image---------------------------------------------------
FROM busybox:glibc

COPY go-api-server /bin/api

EXPOSE 8080
#This command can be overridden by CLI
CMD ["start","-b"]

#This is the point where app starts
ENTRYPOINT ["/bin/api"]


#------------------------Using golang:alpine image---------------------------------------------------
#FROM golang:1.10-alpine3.7
#
#COPY . /go/src/github.com/faem/go-api-server
#WORKDIR /go/src/github.com/faem/go-api-server
#
#RUN go build
#
#CMD ["start","-b"]
#ENTRYPOINT ["/go/src/github.com/faem/go-api-server/go-api-server"]


#------------------------Using Ubuntu image---------------------------------------------------
#FROM ubuntu:18.10
#
#COPY go-api-server /bin/go-api-server
#
#CMD ["start","-b"]
#ENTRYPOINT ["/bin/go-api-server"]


#------------------------Using golang image and bin folder---------------------------------------------------
#FROM golang
#
#ADD . /go/src/github.com/faem/go-api-server
#
#RUN go get github.com/spf13/cobra
#RUN go get github.com/dgrijalva/jwt-go
#RUN go get github.com/gorilla/mux
#RUN go install go-api-server/
#
#ENTRYPOINT ["/go/bin/go-api-server"]


#------------------------Using golang image and src folder---------------------------------------------------
#FROM golang
#
#COPY . /go/src/github.com/faem/go-api-server
#
#RUN go get github.com/spf13/cobra
#RUN go get github.com/dgrijalva/jwt-go
#RUN go get github.com/gorilla/mux
#RUN go build go-api-server/
#
#ENTRYPOINT ["/go/src/github.com/faem/go-api-server/go-api-server"]