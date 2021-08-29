# get image from github
docker pull ghcr.io/wwami-pipeline/nwhcp-server:test

# Run server from docker img
docker run \
--env-file ./docker.env \
-p 80:80 \
--name gateway ghcr.io/wwami-pipeline/nwhcp-server:test;