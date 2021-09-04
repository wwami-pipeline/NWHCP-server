# Run server from docker img
docker run --rm --env-file ./docker.env --network host --name nwhcpgateway ghcr.io/wwami-pipeline/nwhcp-server:development