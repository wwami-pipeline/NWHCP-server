#!/usr/bin/env bash

echo "Building Pipeline Database for Linux..."
docker build -t annaqzhou/nwhcp-server .
go clean

docker push annaqzhou/nwhcp-server

ssh annaz4@v0221.host.s.uw.edu < deploy.sh