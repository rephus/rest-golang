# Simple Golang REST HTTP docker service

## Usage with docker

docker-compose build
docker-compose up

## Endpoint samples 

* POST: curl -X POST  http://localhost:8080/testPOST -d "{\"foo\": \"bar\"}" 
* GET: http://localhost:8080/testGET?foo=bar

[Tutorial](https://thenewstack.io/make-a-restful-json-api-go/)