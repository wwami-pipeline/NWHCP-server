# To run locally
GOOS=linux go build
docker build -t annaqzhou/nwhcp-gateway . # update to your username/nwhcp-gateway
go clean

# Dockerhub
docker push annaqzhou/nwhcp-gateway

# Login to VM and run deploy script
# Jon Schilliing - to get account
ssh annaz4@v0221.host.s.uw.edu < deploy.sh