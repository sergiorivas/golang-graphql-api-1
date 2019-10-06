# Test

## Install

Config bin

gin

sql-migrate

dep

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
