FROM alpine
RUN apk add --no-cache libc6-compat #needed to run go apps on arm64 alpine
COPY NWHCP-server /NWHCP-server
EXPOSE 80
#CMD ["sleep","3600"] #for debugging only
ENTRYPOINT ["/NWHCP-server"]