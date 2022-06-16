FROM golang AS build

WORKDIR /app/whatsapp-signal-bridge

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY * .

RUN go build -o /bot

FROM ubuntu:jammy

WORKDIR /app
COPY --from=build /app/bot /app/bot

# Install Dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    wget \
    openjdk-17-jdk

# Install Signal CLI
ENV VERSION=0.10.8
RUN wget https://github.com/AsamK/signal-cli/releases/download/v"${VERSION}"/signal-cli-"${VERSION}"-Linux.tar.gz && \
    tar xf signal-cli-"${VERSION}"-Linux.tar.gz -C /opt && \
    ln -sf /opt/signal-cli-"${VERSION}"/bin/signal-cli /usr/local/bin/


