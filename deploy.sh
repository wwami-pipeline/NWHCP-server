#!/usr/bin/env bash

./build.sh

echo "Deploying server to dockerhub"
docker build -t andrewk7/pipeline-microservice .
docker push andrewk7/pipeline-microservice

echo "Removing current pipelineMicroservice container"
docker rm -f pipelineMicroservice

./run.sh

go clean
