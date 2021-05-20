FROM golang:alpine as builder

WORKDIR /organizations

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=mod -o organizations

# FROM alpine:3.8
FROM scratch

WORKDIR /root/

COPY --from=builder /organizations .

EXPOSE 5000

EXPOSE 4005

ENV APP_ENV production

CMD ["./organizations"]