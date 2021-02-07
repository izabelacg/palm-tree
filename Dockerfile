FROM ubuntu
WORKDIR /app

RUN apt-get update && apt-get install -y ruby && \
    apt-get -y install jq && \
    gem install wavefront-cli