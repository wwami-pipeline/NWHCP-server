# Run server from docker img
docker run -d --rm --env-file ./docker.env --network host --name nwhcpgateway ghcr.io/wwami-pipeline/nwhcp-server:test-amd64