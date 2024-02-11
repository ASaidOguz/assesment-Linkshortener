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
|               |   |   └── http-handler_test.go
|               |   |   └── http-handler-imp.go
|               |   |   └── http-handler.go
|               |   └── interfaces
|               |   |   └── shortener-service-imp.go
|               |   └── mapper
|               |   └── services
|               |       └── shortener-servic_test.go
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

- [x]  Currently using In memory for database needs.

- [x]  DDD principles're used.

- [x]  Simple Url validation's added. 

- [x]  Tests're added.


## Installation

1. Clone the repository:

   ```
   git clone https://github.com/ASaidOguz/assesment-Linkshortener.git

   ```

2. Navigate to the project directory:
```
cd assesment-Linkshortener

```

3. Install dependencies(test uses mock by go-team)

```
go mod tidy

```

4. Run the project via 

```
go run cmd/marketplace/main.go

```

## Usage

To shorten a URL, send a POST request to http://localhost:8080/shorten with the url parameter set to the original URL. For example:

```
curl -X POST -d 'url=https://example.com' http://localhost:8080/shorten

```

## Configuration
The project currently uses in-memory storage for database needs. No additional configuration is required.

## Testing
To run tests for the project, execute the following command:
Service tests(using -v for gaining more information about tests)

```
go test ./internal/application/services -v
```

Handler tests

```
go test ./internal/application/handlers -v
```

or Test's can be called selectively by using function-selector so can be inspected much better as shown below

Example:

```
go test ./internal/application/handlers -run TestRedirectURLHandler_Works -v

```