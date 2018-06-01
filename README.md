# Waypoint Database


## About
Database runs on MongoDB allows for getting by ID/name, inserting, updating, and deleting of schools and organizations.

### Future Work
- Filtering could be done on database side in order to speed up process 
    - Distance
    - Name
- Currently filter takes on JSON, could recieve CSV file instead of having API translate CSV into JSON.

## Tech and Dependencies
MongoDB and Docker are used to containerize this microservice, which is deployed onto nwhealthcareerpath server.

## Setup
Clone down the repository. 

To run locally:
  - Navigate to https://hub.docker.com and create own dockerhub that will be deployed to
  - Change username and docker container name in deploy.sh script.
  - Run ./deploy.sh script
  
To run on server:
  - Log into server nwhealthcareerpath.uw.edu
  - `docker create network pipeline`
  - Start MongoDB instance with 
    ```
    docker run -d \
    -p 27017:27017 \
    --network pipeline \
    --name pipelinedb \
    mongo
    ```
  - Start pipelinemicroservice with: 
    ```
    docker run -d -p 127.0.0.1:4002:4002 \
    --network pipeline \
    --name pipelinemicroservice \
    -e DBADDR=pipelinedb:27017 \
    -e PORTADDR=:4002 \
    YOUR_DOCKER_USERNAME_HERE/YOUR_DOCKER_CONTAINER_NAME_HERE
    ```
  
