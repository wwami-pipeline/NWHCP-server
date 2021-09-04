# compile
cd gateway
GOOS=linux go build
cd ..

# make docker
docker build -t loibucket/nwhcp-gateway . # update to your username/nwhcp-gateway
docker rm -f nwhcp-gateway || true # mostly for testing because have to remove docker img every time you redeploy

# Run server from docker img
docker run \
--rm \
--network host \
--env-file ./docker.env \
--name nwhcpgateway loibucket/nwhcp-gateway;
