FROM mongo
WORKDIR /root
COPY . ./json
RUN chmod +x ./json/import.sh