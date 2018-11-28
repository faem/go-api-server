#---------------------------------------------------------------------------
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


#---------------------------------------------------------------------------
FROM ubuntu:18.10

COPY LinkedinApiServer /bin/LinkedinApiServer

ENTRYPOINT ["/bin/LinkedinApiServer"]



#---------------------------------------------------------------------------
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