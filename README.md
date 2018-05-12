# Gin based  WebApp Project


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

- [] Database/ORM
- [] Middleware
- [] Test
- [] Cache/Session
