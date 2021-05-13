GOOS=linux go build
docker build -t annaqzhou/summary .
go clean

docker push annaqzhou/summary

ssh ec2-user@ec2-13-52-124-193.us-west-1.compute.amazonaws.com < deploy.sh