docker rm -f summary

docker pull annaqzhou/summmary
export ADDR="summary:80"

docker run -d \
--network aznet \
-e ADDR=$ADDR \
--name summary annaqzhou/summary
