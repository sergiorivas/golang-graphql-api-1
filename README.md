This a template for an GraphQL API in Golang with

- [Gin Server](https://github.com/gin-gonic/gin): Fast server
- [Graphql go (Gohpers)](https://github.com/graph-gophers/graphql-go): GraphQL Engine
- [Gorm](https://gorm.io/docs/): ORM in GO (connected with postgreSQL)
- [GraphiQL](https://github.com/graphql/graphiql): GraphQL web client
- [SQL-Migrate](https://github.com/rubenv/sql-migrate): DB migrations
- [Dep](https://github.com/golang/dep): Go dependecy manager

# Install

You can install `asdf` if you are in linux or MacOS, this allow you to have several dev envs including GOLANG

### 1. Config go path and bin
Add in your `~/.bashrc` or `~/.zshrc`

```
export GOPATH=$HOME/go
export PATH="$PATH:$GOPATH/bin"
```

---

Now, you have to install some commands (2. ,3. ,4.)

### 2. Gin (Server with hot-reload)
```
go get github.com/codegangsta/gin
```

### 3. SQL-Migrate (for migrations)
```
go get -v github.com/rubenv/sql-migrate/...
```

### 4. Dep (Dependecy manager)
```
go get -u github.com/golang/dep/cmd/dep
```

---

### 5. Copy .env file
```
cp .env-example .env
```

And change your **configuration values**

# Initial Run (preparing database)

### 1. Get depencies

```
dep ensure
```

### 2. Run initial migrations

```
export `egrep -v "#" .env | xargs -0`
sql-migrate up -config="config/sql-migrate.yml"
```

### 3. Run seeds
```
go run db/seeds/adding_articles.go
```

# Regular Run

### Start the server
```
gin
```

The API should be at `http://localhost:3000/graphql`

# Migrations

Run this at least once in your console

```
export `egrep -v "#" .env | xargs -0`
```

----

Creating a migration
```
sql-migrate new  -config="config/sql-migrate.yml" <name>
```

Migrating
```
sql-migrate up -config="config/sql-migrate.yml"
```

Rollback
```
sql-migrate down -config="config/sql-migrate.yml"
```
