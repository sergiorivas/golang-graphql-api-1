# Test

## Install

You can install `asdf` if you are in linux, this allow you to have several dev envs including GOLANG

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

## Config

```
export `egrep -v "#" .env | xargs -0`
```

## Initial Run

## Regular Run

```
gin
```

Open `http://localhost:3000` for the playground

## Migrations

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
