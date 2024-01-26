FROM alpine:latest
MAINTAINER Juan Diego Castellanos Jerez
LABEL authors="Juan Diego Castellanos Jerez"

RUN apk --no-cache update && apk --no-cache add wget bash openjdk8 curl \
    && rm -rf /var/cache/apk/*


# Me queda pendiente hacer que se exponga el servicio de zookeeper y despues crear el
# container de kafka broker para conectarlos

WORKDIR ./
COPY persistence/kafka-3.6.0-src ./
#RUN cd kafka-3.6.0-src/
RUN chmod +x ./gradlew
RUN ./gradlew jar -PscalaVersion=2.13.11
#RUN chmod +x ./gradlew && ./gradlew jar -PscalaVersion=2.13.11
ENTRYPOINT ["tail", "-f", "/dev/null"]
