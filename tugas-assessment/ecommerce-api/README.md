# Ecommerce API

This program is a REST API for a simple e-commerce website written in Go with the following features:

- Create listings for products to sell
- List available products to buy with filtering
- Add products to a shopping cart and checkout

## Problem & Motivation

A mini-project to create a back-end API for a simple e-commerce website using Go programming languange.

## How to Run

Initialize Go in workspace

```
go mod init
```

Import main dependencies:

- Echo http framework
- PostgreSQL GORM

```
go get github.com/labstack/echo/v4
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

To run program locally, uncomment `"DB_HOST=127.0.0.1"` and comment `"DB_HOST=ecommerce-postgres"` in `.env` file then run command:

```
go run main.go
```

To run program in docker compose, revert previous comments in `.env` file then run commands:

```
docker-compose build
docker-compose up -d
```

To stop the program running in docker containers, run:

```
docker-compose down
```

## API Documentation

https://docs.google.com/spreadsheets/d/1m7vjL-PufRdo4NbXtVnx8B70KasnpftmWPlb7RrH3hE/edit?usp=sharing
