#!/usr/bin/env bash
docker pull andrewk7/pipeline-microservice

# make network on server
# docker network create pipeline

# echo "Running MongoDB"

docker run -d \
-p 27017:27017 \
--network pipeline \
--name pipelinedb \
mongo

echo "Running pipeline-microservice"

# docker run -d -p 443:443 \
# --network pipeline \
# --name pipelineMicroservice \
# -v /etc/pki/tls/:/tls/:ro \
# -e DBADDR=pipelinedb:27017 \
# -e TLSCERT=/tls/certs/nwhealthcareerpath.uw.edu.crt \
# -e TLSKEY=/tls/private/nwhealthcareerpath.uw.edu.key \
# andrewk7/pipeline-microservice

docker run -d -p 4002:4002 \
--network pipeline \
--name pipelinemicroservice \
-e DBADDR=pipelinedb:27017 \
-e PORTADDR=:4002 \
andrewk7/pipeline-microservice