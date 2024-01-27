FROM alpine:latest
MAINTAINER Juan Diego Castellanos Jerez
LABEL authors="Juan Diego Castellanos Jerez"

RUN apk --no-cache update && apk --no-cache add wget bash openjdk8 curl \
    && rm -rf /var/cache/apk/*

WORKDIR ./
COPY persistence/kafka_2.13-3.6.1 ./
RUN  KAFKA_CLUSTER_ID=$(bin/kafka-storage.sh random-uuid) && \
     bin/kafka-storage.sh format --config config/kraft/server.properties --cluster-id $KAFKA_CLUSTER_ID
#CMD ./bin/kafka-storage.sh format -t $KAFKA_CLUSTER_ID -c config/kraft/server.properties
#CMD ["bin/kafka-storage.sh format", "--config config/kraft/server.properties", "--cluster-id $KAFKA_CLUSTER_ID"]
CMD ["bin/kafka-server-start.sh", "config/kraft/server.properties"]
#ENTRYPOINT ["tail", "-f", "/dev/null"]
EXPOSE 9092