# CURL commands

```
curl -s -X GET <link>
curl -s -X GET <link>
curl -s -X DELETE <link>
curl -s -X POST -H 'Content-Type: application/json' -d '<json data>' <link>
curl -s -X PUT -H 'Content-Type: application/json' -d '<json data>' <link>
curl --user <username>:<name> <link>
```

## Testing

##### Read all profiles
```
curl -s -X GET localhost:8080/in | jq
```
##### Read a profile
```
curl -s -X GET localhost:8080/in/masud-rahman | jq
```
##### Delete a profile
```
curl -s -X DELETE localhost:8080/in/masud-rahman | jq
```
##### Create a profile
```
curl -s -X POST -H 'Content-Type: application/json' -d '{"id":"kfoozminus","name":"Jannatul Ferdous","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C++","noOfEndorsement":100},{"name":"C","noOfEndorsement":100}]}' localhost:8080/in | jq
```
##### Update a profile info
```
curl -s -X PUT -H 'Content-Type: application/json' -d '{"id":"masud-rahman","name":"Masudur Rahman (modified)","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C","noOfEndorsement":3},{"name":"C++","noOfEndorsement":4}]}' localhost:8080/in/masud-rahman | jq
```
##### Valid Authorization check
```
curl --user fahim:1234 localhost:8080/in
```
##### Invalid Authorization Header check
```
curl --user fahim:dfsd:d localhost:8080/in
```
#### Authorize
```
curl -H 'Accept: application/json' -H "Authorization: Bearer $eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTQzMjU5OTQ2LCJ1c2VyIjoiZmFoaW0ifQ.qRTYLq4en4MMRZdNs3XjhOAOHSrkt_UqZM-xmpnoXIo" https://localhost/in
```