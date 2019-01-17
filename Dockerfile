FROM golang:alpine as builder

WORKDIR /pipeline-db

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o app 

FROM alpine:3.8

WORKDIR /root/

COPY --from=builder /pipeline-db/app .

EXPOSE 4002

CMD ["./app"]