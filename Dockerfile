FROM ubuntu
WORKDIR /app
ADD . /app/
RUN apt-get update && apt-get install -y ruby && \
    gem install wavefront-cli