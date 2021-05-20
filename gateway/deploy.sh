export MYSQL_ROOT_PASSWORD="password"
export MYSQL_DATABASE=mydatabase
export TLSCERT=/etc/letsencrypt/live/nwhealthcareerpath.uw.edu/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/nwhealthcareerpath.uw.edu/privkey.pem
export REDISADDR=myredis:6379
export SESSIONKEY="key"
export SERVER2ADDR="http://organizations:5000"

# docker rm -f helloservertest;


# docker network create verdancynet;

# docker run -d --name myredis --network verdancynet redis;
# docker run --name myredis -d redis;

# docker rm -f verdancy_db;

# docker pull annaqzhou/verdancydb;
# docker run -d \
# -p 3306:3306 \
# --name verdancy_db \
# --network verdancynet \
# -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
# -e MYSQL_DATABASE=$MYSQL_DATABASE \
# annaqzhou/verdancydb;

docker rm -f gateway;
docker pull annaqzhou/nwhcp-gateway;

docker run -d -p 443:443 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
-e SESSIONKEY=$SESSIONKEY \
-e REDISADDR=$REDISADDR \
-e SUMMARYADDR=$SUMMARYADDR \
-e DSN=root:$MYSQL_ROOT_PASSWORD@tcp\(nwhcp-sqldb:3306\)/$MYSQL_DATABASE \
-e SERVER2ADDR=$SERVER2ADDR \
--network nwhcp-docker_default \
--name gateway annaqzhou/nwhcp-gateway;