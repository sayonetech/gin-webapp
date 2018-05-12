# Gin based REST API Project


#### Environment Requirements

- GO >= 1.8

### Install

```
cd $GOPATH/src

git clone https://github.com/sayonetech/gin-webapp.git

```
### Load Dependency

```
cd gin-webapp
make deps
```

#### Build Service
```
make build
```

#### Run the Service
```
make
```

visit by browser: http://localhost:4000/api/index

#### Database Migration
```
make migrate
```
## TODO

- [x] Database/ORM
- [x] Migration
- [] Authentication
- [] Session
- [] ElasticSearch
- [] Task Queue
- [] SMTP
- [] Middleware
- [] GZip https://github.com/gin-contrib/gzip
- [] Test
- [] Cache


## Task Q
  * https://github.com/gocelery/gocelery
  * https://github.com/RichardKnop/machinery
  * https://eng.uber.com/cherami/

## Email
 * https://github.com/go-gomail/gomail
