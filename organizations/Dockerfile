FROM golang:alpine as builder

WORKDIR /organizations

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o NWHCP-server

# FROM alpine:3.8
FROM scratch

WORKDIR /root/

COPY --from=builder /organizations/organizations .

ENV APP_ENV production

CMD ["./organizations"]