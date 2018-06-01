#!/usr/bin/env bash

### REPLACE WITH OWN DOCKERHUB ACCOUNT AND CONTAINER NAME
# docker pull andrewk7/pipeline-microservice

# make network on server
# docker network create pipeline

echo "Running MongoDB"

docker run -d \
-p 27017:27017 \
--network pipeline \
--name pipelinedb \
mongo

echo "Running pipeline-microservice"

### REPLACE WITH OWN DOCKERHUB ACCOUNT AND CONTAINER NAME

# docker run -d -p 127.0.0.1:4002:4002 \
# --network pipeline \
# --name pipelinemicroservice \
# -e DBADDR=pipelinedb:27017 \
# -e PORTADDR=:4002 \
# andrewk7/pipeline-microservice