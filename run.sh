#!/usr/bin/env bash
docker pull andrewk7/pipeline-microservice

# make network on server
# docker network create pipeline

# echo "Running MongoDB"

# docker run -d \
# -p 27017:27017 \
# --network pipeline \
# --name pipelinedb \
# mongo

echo "Running pipeline-microservice"

docker run -d -p 4000:4000 \
--network pipeline \
--name pipelineMicroservice \
-e DBADDR=pipelinedb:27017 \
andrewk7/pipeline-microservice
