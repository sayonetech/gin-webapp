#!/bin/sh
## intiate docker container
#docker rm -f go-webapp
docker run -it --name=go-webapp --hostname=dockerGo --env-file=env/docker-local.env $1 /bin/bash
