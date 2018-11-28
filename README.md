# HTTP Api Server
This repo contains an API server which is built using Golang. It uses Gorilla mux library for HTTP API calls,Cobra CLI for making a CLI of the application and go-jwt for adding jwt authorization token.

Here GET request is implemented using Basic AuthN and POST, PUT and DELETE requests are implemented using JWT AuthZ. 

## Run directly from Docker-Hub
```
docker run -p 8080:8080 fahimabrar/api start
```
After running this use the curl commands to test the server

## Run using dockerfile
```
docker build -t <give a name> .
docker run -p 8080:8080 <given name> start
```

### Upload to Docker-Hub
```
docker login --username=<docker hub username>
docker tag <id of the created image> <docker hub username>/<name of the image>:<tag>
docker push <docker hub username>/<name of the image>:<tag>
```

## Run from Source Code
```
go build
```

##### To start the server with default configuration

```
./LinkedinApiServer start
```
##### To give a port number
```
./LinkedinApiServer start -p <port number>
```
##### To bypass login information
```
./LinkedinApiServer start -b
```
##### To stop the server after a definite time
``` 
./LinkedinApiServer start -s <minute>
```
##### To check the version of the API
```
./LinkedinApiServer version
```
##### To get a JWT token
```
./LinkedinApiServer gentkn -u <username> -e <expiration time in minute>
```

## CURL commands

##### Read all profiles
```
curl --user admin:admin -s -X GET localhost:8080/in
```
##### Read a profile
```
curl --user admin:admin -s -X GET localhost:8080/in/fahim-abrar
```
##### Delete a profile
```
curl -H "Authorization: Bearer <token>" -s -X DELETE localhost:8080/in/masud-rahman
```
##### Create a profile
```
curl -H "Authorization: Bearer <token>" -s -X POST -H 'Content-Type: application/json' -d '{"id":"kfoozminus","name":"Jannatul Ferdous","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C++","noOfEndorsement":100},{"name":"C","noOfEndorsement":100}]}' localhost:8080/in
```
##### Update a profile info
```
curl -H "Authorization: Bearer <token>" -s -X PUT -H 'Content-Type: application/json' -d '{"id":"masud-rahman","name":"Masudur Rahman (modified)","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C","noOfEndorsement":3},{"name":"C++","noOfEndorsement":4}]}' localhost:8080/in/masud-rahman
```
##### Generate a JWT token
```
curl --user admin:admin localhost:8080/token/<expiration time in minute>
```
##### Invalid Authorization Header check
```
curl --user fahim:dfsd:d localhost:8080/in
```
##### Expired Token Check
```
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTQzMjU5OTQ2LCJ1c2VyIjoiZmFoaW0ifQ.qRTYLq4en4MMRZdNs3XjhOAOHSrkt_UqZM-xmpnoXIo" -s -X POST -H 'Content-Type: application/json' -d '{"id":"kfoozminus","name":"Jannatul Ferdous","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C++","noOfEndorsement":100},{"name":"C","noOfEndorsement":100}]}' localhost:8080/in
```
##### Shutdown the Server
``` 
curl localhost:8080/shutdown
```