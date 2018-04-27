FROM alpine
COPY pipeline-db /pipeline-db
EXPOSE 80
ENTRYPOINT ["/pipeline-db"]
