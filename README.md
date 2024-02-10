# Url-Shortener with Domain-Driven-Design

This is assesment repository for Url-shortener made using only standard libraries and taking account of domain-driven-design principles and structure patterns.

Project structure is 

```

├── github.com
│   └── ASaidOguz
│       └── assesment-Linkshortener
│           |── cmd
│           |    └── marketplace
|           |        └──main.go
│           └── internal
|               └── application
|               |   └── command
|               |   └── handlers
|               |   |   └── http-handler.go
|               |   └── interfaces
|               |   |   └── shortener-service.go
|               |   └── mapper
|               |   └── services
|               |       └── shortener-service-imp.go
|               |       └── shortener-service.go
|               └── config
|               |   └── config.go
|               └── domain
|               |   └── entity
|               |   |   └── shortenedurl.go
|               |   |   └── url.go
|               |   └── repositories
|               |       └── repository.go
|               └── infrastructure
|               |   └── inmemory.go
|               └── interface

```

After starting ; project listens port:8080 and parsing urls as it received as post request with form values 

key: url            // required key for POST request.

value: original url // that requested to be shortened. 

 [x] Currently using In memory for database needs 