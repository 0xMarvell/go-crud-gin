# Simple CRUD API (WORK IN PROGRESS ⚠️)

A simple blog API built using:

- [Gin-Gonic](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)

## DOCS

Work in progress..

## Environment Variables

This project makes use of the [viper](https://github.com/spf13/viper) package to store environment variables. To run this project on your local machine, you will need to add your Postgres database details as environment variables. This [article](https://blog.logrocket.com/handling-go-configuration-viper/) explains how to use the package.
> An `.env.example` file has been provided to help with setting up your environment variables.

## Run API locally

- Clone Repo

    ```bash
    $ git clone https://github.com/0xMarvell/simple-blog-posts.git
    ```

- Make sure to have [Go](https://go.dev/) installed on your local machine
- Open the code base directory in terminal
- Launch API server:

    ```go
    $ go build -o blogposts cmd/main.go
    $ ./blogposts
    ```

> Test API using Postman
