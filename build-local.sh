# compile
cd gateway
GOOS=linux go build
cd ..

# make docker
docker build -t localbuild/nwhcp-gateway . # update to your username/nwhcp-gateway
docker rm -f nwhcpgateway || true # mostly for testing because have to remove docker img every time you redeploy

# Run server from docker img
docker run \
--rm \
--network host \
--env-file ./docker.env \
--name nwhcpgateway localbuild/nwhcp-gateway;
