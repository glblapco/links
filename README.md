# Links

Links archive using microservices architecture.
For now it doesn't store data in a database, you can store your links locally piping the response to a file.

![links demo](/static/linksDemo.png)

## Usage.

Why make GUIs when you can just use `cURL` and `jq`?

### With Docker

1. Clone this repository.
2. `$ docker build -t b-ap/links:1.0 .`
3. `$ docker images | grep links`
4. `$ docker run -p 80:80 b-ap/links:1.0`
5. `$ curl -X POST -H "Content-Type: application/json" -d '{"title": "b-ap website", "url": "b-ap.xyz"}' localhost:8000`
6. You can also use `jq` to format the response:
`$ curl -X POST -H "Content-Type: application/json" -d '{"title": "b-ap website", "url": "b-ap.xyz"}' localhost:8000 | jq`
7. To list all your links, just send a `GET` request to the page:
`$ curl -X GET localhost:8000 | jq`
8. If you want to store it locally, just pipe the response to a file:
`$ curl -X GET localhost:8000 | jq > links.txt`

### Without Docker.

1. Clone this repository.
2. `$ go run links.go`
3. `$ curl -X POST -H "Content-Type: application/json" -d '{"title": "b-ap website", "url": "b-ap.xyz"}' localhost:8000`
4. You can also use `jq` to format the response:
`$ curl -X POST -H "Content-Type: application/json" -d '{"title": "b-ap website", "url": "b-ap.xyz"}' localhost:8000 | jq`
5. To list all your links, just send a `GET` request to the page:
`$ curl -X GET localhost:8000 | jq`
6. If you want to store it locally, just pipe the response to a file:
`$ curl -X GET localhost:8000 | jq > links.txt`
