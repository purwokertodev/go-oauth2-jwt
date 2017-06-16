OAUTH2 example with GO, JWT-go and gorilla

[![Build Status](https://travis-ci.org/wuriyanto48/go-oauth2-jwt.svg?branch=master)](https://travis-ci.org/wuriyanto48/go-oauth2-jwt)

before get started, you have to know :

- https://golang.org/doc/code.html

# Get Started
- Install Glide https://github.com/Masterminds/glide
- Install dependencies
  ```language:shell
  $ glide install
  ```

after this step, you can do this:
- go run main.go

try it with curl:

```language:shell
curl localhost:9000/token?grant_type=password -X POST -d "{"username": "wuriyanto", "password": "123456"}'
```

if you using windows:
```language:shell
curl localhost:9000/token?grant_type=password -X POST -d "{\"username\": \"wuriyanto\", \"password\": \"123456\"}'
```

# To Do
- [x] Password Grant Type
- [ ] Add Client Credentials Grant
- [ ] Add new Code
- [ ] Writing test



***Copyright 2017 Wuriyanto Musobar***
