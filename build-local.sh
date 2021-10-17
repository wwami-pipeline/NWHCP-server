# compile
cd gateway
GOOS=linux go build
cd ..

# make docker
docker build --no-cache -t localbuild/nwhcp-gateway . # update to your username/nwhcp-gateway
docker rm -f nwhcpgateway || true # mostly for testing because have to remove docker img every time you redeploy

# Run server from docker img
# host.docker.internal:6379 needed for mac users
docker run \
--rm \
--publish 8080:8080 \
--network dockernet \
--env REDIS_ADDR="host.docker.internal:6379" \
--name nwhcpgateway localbuild/nwhcp-gateway
