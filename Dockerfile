FROM ubuntu:latest

ENV HTTP_HOST "0.0.0.0"
ENV HTTP_PORT "9090"
ENV LOG_LEVEL "info"
ENV LOG_FORMAT "text"

RUN apt-get update && apt-get upgrade -y --no-install-recommends
RUN apt-get install -y ca-certificates
RUN apt-get install -y dumb-init

ADD ./miningpost_linux_amd64 /usr/bin/miningpost

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD miningpost