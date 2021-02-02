FROM golang:alpine as builder

WORKDIR /NWHCP-server

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o NWHCP-server

# FROM alpine:3.8
FROM scratch

WORKDIR /root/

COPY --from=builder /NWHCP-server/NWHCP-server .

ENV APP_ENV production

CMD ["./NWHCP-server"]