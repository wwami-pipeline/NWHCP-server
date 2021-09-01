# compile
cd gateway
GOOS=linux go build
cd ..

# make docker
docker build -t loibucket/nwhcp-gateway . # update to your username/nwhcp-gateway
docker rm -f gateway || true # mostly for testing because have to remove docker img every time you redeploy

# Run server from docker img
docker run \
--env-file ./docker.env \
-p 80:80 \
--name gateway loibucket/nwhcp-gateway;