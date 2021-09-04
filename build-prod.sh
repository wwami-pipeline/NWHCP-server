# run on webserver
docker run -d --network host --rm --pull=always --name nwhcpgateway ghcr.io/wwami-pipeline/nwhcp-server:test-amd64
# docker run --network host --rm --pull=always --name nwhcpgateway ghcr.io/wwami-pipeline/nwhcp-server:test-amd64