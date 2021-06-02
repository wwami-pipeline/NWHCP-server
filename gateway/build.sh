GOOS=linux go build
docker build -t annaqzhou/nwhcp-gateway .
go clean

docker push annaqzhou/nwhcp-gateway

ssh annaz4@v0221.host.s.uw.edu < deploy.sh