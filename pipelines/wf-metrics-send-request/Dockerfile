FROM ubuntu as base
WORKDIR /app

RUN apt-get update && \
    apt-get install -y siege && \
    apt-get install -y curl && \
    apt-get install -y git make zip build-essential libssl-dev

FROM base

RUN git clone https://github.com/wg/wrk.git && cd wrk
RUN cd wrk && \
    make && cp wrk /usr/local/bin