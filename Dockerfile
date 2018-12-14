#------------------------Using busybox image---------------------------------------------------
FROM busybox:glibc

COPY LinkedinApiServer /bin/api

EXPOSE 8080
#This command can be overridden by CLI
CMD ["start","-b"]

#This is the point where app starts
ENTRYPOINT ["/bin/api"]


#------------------------Using golang:alpine image---------------------------------------------------
#FROM golang:1.10-alpine3.7
#
#COPY . /go/src/LinkedinApiServer
#WORKDIR /go/src/LinkedinApiServer
#
#RUN go build
#
#CMD ["start","-b"]
#ENTRYPOINT ["/go/src/LinkedinApiServer/LinkedinApiServer"]


#------------------------Using Ubuntu image---------------------------------------------------
#FROM ubuntu:18.10
#
#COPY LinkedinApiServer /bin/LinkedinApiServer
#
#CMD ["start","-b"]
#ENTRYPOINT ["/bin/LinkedinApiServer"]


#------------------------Using golang image and bin folder---------------------------------------------------
#FROM golang
#
#ADD . /go/src/LinkedinApiServer
#
#RUN go get github.com/spf13/cobra
#RUN go get github.com/dgrijalva/jwt-go
#RUN go get github.com/gorilla/mux
#RUN go install LinkedinApiServer/
#
#ENTRYPOINT ["/go/bin/LinkedinApiServer"]


#------------------------Using golang image and src folder---------------------------------------------------
#FROM golang
#
#COPY . /go/src/LinkedinApiServer
#
#RUN go get github.com/spf13/cobra
#RUN go get github.com/dgrijalva/jwt-go
#RUN go get github.com/gorilla/mux
#RUN go build LinkedinApiServer/
#
#ENTRYPOINT ["/go/src/LinkedinApiServer/LinkedinApiServer"]