# RSS Aggregator

A backend server in Golang for aggregating RSS feeds from the internet. This project uses PostgreSQL as the database and employs sqlc for managing SQL queries and goose for database migrations.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Clone the Repository](#clone-the-repository)
- [Setup PostgreSQL](#setup-postgresql)
- [Install Dependencies](#install-dependencies)
- [Database Migrations](#database-migrations)
- [Running the Project](#running-the-project)


## Features
- Aggregates RSS feeds from various sources
- Stores aggregated data in PostgreSQL
- Provides a RESTful API to interact with the aggregated data

## Prerequisites
- Go 1.16 or later
- PostgreSQL 13 or later

## Installation
### Clone the Repository
```
git clone https://github.com/ayushrakesh/go-rssagg.git
```
```
cd go-rssagg
```

### Setup PostgreSQL
1. Install PostgreSQL:

**On macOS:**
```
brew install postgresql
```

**On Ubuntu:**
```
sudo apt update
```
```
sudo apt install postgresql postgresql-contrib
```

2. Start PostgreSQL service:
```
sudo service postgresql start
```

3. Create a new database and user:
```
CREATE DATABASE rssagg;
CREATE USER rssuser WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE rssagg TO rssuser;
```

### Install Dependencies

1. Install goose for database migrations:
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

3. Install sqlc for SQL query management:
```
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

### Database Migrations

1. Apply the database migrations:
```
goose -dir ./sql/schema postgres "user=rssuser password=password dbname=rssagg sslmode=disable" up
```

## Running the Project

1. Generate the SQL code with sqlc:
```
sqlc generate
```
2. Run the project:
```
go run main.go
```

