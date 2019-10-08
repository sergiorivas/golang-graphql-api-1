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

### 1. Run initial migrations

```
export `egrep -v "#" .env | xargs -0`
sql-migrate up -config="config/sql-migrate.yml"
```

### 2. Run seeds
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
