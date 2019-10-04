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

1. Args
2. Mutation
3. Database migration
4. Database seeds
5. Add dataloader
6.  Clean the project

### Done

