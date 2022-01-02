# Links

Links archive using microservices architecture.
For now it doesn't store data in a database, you can store your links locally piping the response to a file.

![links demo](/static/linksDemo.png)

## Usage.

Why make GUIs when you can just use `cURL` and `jq`?

### Without docker.

1. `$ go run links.go`
2. `$ curl -X POST -H "Content-Type: application/json" -d '{"title": "b-ap website", "url": "b-ap.xyz"}' localhost:8000`
3. You can also use `jq` to format the response:
`$ curl -X POST -H "Content-Type: application/json" -d '{"title": "b-ap website", "url": "b-ap.xyz"}' localhost:8000 | jq`
4. To list all your links, just send a `GET` request to the page:
`$ curl -X GET localhost:8000 | jq`
5. If you want to store it locally, just pipe the response to a file:
`$ curl -X GET localhost:8000 | jq > links.txt`
