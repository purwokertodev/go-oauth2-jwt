OAUTH2 example with GO, JWT-go and gorilla

before get started, you have to know :

- https://golang.org/doc/code.html

after this step, you can do this:
- go run main.go

try it with curl:

```language:shell
curl localhost:3000/token?grant_type=password -X POST -d "{"username": "wuriyanto", "password": "123456"}'
```

if you using windows:
```language:shell
curl localhost:3000/token?grant_type=password -X POST -d "{\"username\": \"wuriyanto\", \"password\": \"123456\"}'
```