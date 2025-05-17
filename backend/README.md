# Golang

## Create module
```
go mod init certificate-ledger
```

## Install library

### Library create UUID
```
go get github.com/google/uuid
```
### Router HTTP
```
go get github.com/gorilla/mux
```
### Database (Driver) MySQL
```
go get github.com/go-sql-driver/mysql
```
### Additional encryption algorithm
```
go get golang.org/x/crypto
```
### Library to load environment variables from .env file
```
go get github.com/joho/godotenv
```
### JSON Web Token (JWT) Generation and Validation Library
```
go get github.com/golang-jwt/jwt/v5
```
## Run server
```
go run cmd/main.go
```