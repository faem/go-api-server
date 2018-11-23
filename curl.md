# CURL commands

```
curl -s -X GET <link>
curl -s -X GET <link>
curl -s -X DELETE <link>
curl -s -X POST -H 'Content-Type: application/json' -d '<json data>' <link>
curl -s -X PUT -H 'Content-Type: application/json' -d '<json data>' <link>
```

## Demo Curl commands

```
curl -s -X GET localhost:8080/in | jq
curl -s -X GET localhost:8080/in/masud-rahman | jq
curl -s -X DELETE localhost:8080/in/masud-rahman | jq
curl -s -X POST -H 'Content-Type: application/json' -d '{"id":"kfoozminus","name":"Jannatul Ferdous","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C++","noOfEndorsement":100},{"name":"C","noOfEndorsement":100}]}' localhost:8080/in | jq
curl -s -X PUT -H 'Content-Type: application/json' -d '{"id":"masud-rahman","name":"Masudur Rahman (modified)","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C","noOfEndorsement":3},{"name":"C++","noOfEndorsement":4}]}' localhost:8080/in/masud-rahman | jq
```