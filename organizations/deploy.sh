export PORTADDR="organizations:5000"
export INTERNAL_PORT="organizations:4005"
export DBADDR=nwhcp-mongo:27017
export MYSQL_ROOT_PASSWORD="password"
export MYSQL_DATABASE=mydatabase
export DSN=root:$MYSQL_ROOT_PASSWORD@tcp\(nwhcp-sqldb:3306\)/$MYSQL_DATABASE \

docker rm -f organizations

docker pull annaqzhou/nwhcp-server


docker run -d --network nwhcp-docker_default \
-e PORTADDR=$PORTADDR \
-e INTERNAL_PORT=$INTERNAL_PORT \
-e DBADDR=$DBADDR \
-e DSN=$DSN \
--name organizations annaqzhou/nwhcp-server
