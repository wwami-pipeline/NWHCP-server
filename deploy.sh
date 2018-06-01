#!/usr/bin/env bash

./build.sh

echo "Deploying server to dockerhub"


### REPLACE WITH OWN DOCKERHUB ACCOUNT AND NAME
# docker build -t andrewk7/pipeline-microservice .
# docker push andrewk7/pipeline-microservice



echo "Removing current pipelineMicroservice container"
docker rm -f pipelinemicroservice

./run.sh

go clean
