# Databases Sync Service

## Setup

Dependencies:

- GO 1.24
- mysql <= 8.4
- mysqldump

## Install dependencies

```bash
go mod tidy
```

## Setup environment

1. Create a `db.json` file in the root directory of the project and copy the contents of the `db.example.json` file into it.

## Run app

```bash
go run cmd/main.go
```