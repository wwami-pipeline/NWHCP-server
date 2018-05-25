FROM alpine
RUN apk add --no-cache ca-certificates
COPY pipeline-db /pipeline-db
EXPOSE 4002
ENTRYPOINT ["/pipeline-db"]
