# GopherCorn - Simple GraphQL API written in Golang

[![Go 1.16](https://img.shields.io/badge/Go-v1.16-blue)](https://golang.org/doc/go1.16)

```yaml
Author: Mitch Murphy
Date: 2021 Jun 20
```

![GopherCorn](media/gophergraph.jpeg)


## Notes

Unless you want to manually add implementation to `graph/generted/generated.go`, run the following:

`go run github.com/99designs/gqlgen generate`

## Run

You can run this project locally using the following command: `go run server.go`

### Docker

It is recommended that you run this application using the provided multi-stage Docker file:

```shell
docker build -t gophercorn:0.1.0 .
docker run -p 8080:8080 -d gophercorn:0.1.0
```

Using either approach the application is now accessible at: `http://localhost:8080/`

### Example Query/Mutation

```javascript
mutation createUser {
  createUser(input:{
    name:"Mitch"
    address:"123 Main St"
    phoneNumber:"123-456-7890"
  }) {
    id
    name
    address
    phoneNumber
  }
}

mutation createPost {
  createPost(input:{
    title:"Test Post", 
    body:"Test Body", 
    userId:"T5577006791947779410"
  }) {
    user {
      id
    }
    title
    body
  }
}

query findPosts {
    posts {
      title
      body
      user {
        name
      }
    }
}

query getUser {
  user(userId:"T5577006791947779410") {
    id
    name
    address
    phoneNumber
  }
}

query getUsers {
  users {
    id
    name
    address
    phoneNumber
  }
}
```