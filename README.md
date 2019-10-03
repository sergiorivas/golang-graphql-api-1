# Test

## Generate
```
rm gql/resolvers/generated.go; gqlgen -c config/gqlgen.yml -v
```

See the file `gql/resolvers/generated.go`, Make sure you implement all functions
```
type queryResolver struct{ *Resolver }

... // <-- these functions


```

Delete the file `gql/resolvers/generated.go`

## Run

```
gin run server.go
```

Open `http://localhost:3000` for the playground

---

### Pending

1. Basic Graphql for One Comment
2. Basic Graphql for un Article with comments
3. Bind Database
3. Database migration
4. Database seeds
4. Add Comments Migration
5. Add Comment Migration
6. Add Comment Resolver
7. Add dataloader
8.  Add mutation
9.  Clean the project

### Done

