# GopherCorn - Simple GraphQL API written in Golang

```yaml
Author: Mitch Murphy
Date: 2021 Jun 20
```

**Run**: `go run server.go`

## Example Query/Mutation

```javascript
mutation createTodo {
  createTodo(input:{text:"todo", userId:"1"}) {
    user {
      id
    }
    text
    done
  }
}

query findTodos {
    todos {
      text
      done
      user {
        name
      }
    }
}
```