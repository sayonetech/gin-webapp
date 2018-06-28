#!/bin/sh
docker rm -f go-webapp
docker run -d --name=go-webapp -hostname=dockerGo  --env-file=env/docker-local.env $1 /go-webapp
