# nats-cli [![Build Status](https://travis-ci.org/shadiakiki1986/nats-cli.svg?branch=master)](https://travis-ci.org/shadiakiki1986/nats-cli)
NATS client CLI

Equivalent to the [ruby-nats](https://github.com/nats-io/ruby-nats) example scripts

## Usage

Download the binary from one of the [releases](https://github.com/shadiakiki1986/nats-cli/releases)
(in example below, `amd64` is the output of `dpkg --print-architecture`)
and run

```bash
wget https://github.com/shadiakiki1986/nats-cli/releases/download/0.0.1/nats-amd64 -O /sbin/nats
chmod +x /sbin/nats

# publish to channel "foo" the message "help me!"
bin/nats pub foo "help me!"

# same on different server
bin/nats --server nats://someserver:4222 pub foo "help me!"

# subscribe to channel "foo"
bin/nats sub foo
```

## Development
Pre-requisites
```bash
sudo apt-get install golang
export GOPATH=${PWD}
go get github.com/nats-io/go-nats
go get gopkg.in/urfave/cli.v2
```

Build binary (copied from [gosu](https://github.com/tianon/gosu/blob/master/Dockerfile))

```bash
CGO_ENABLED=0 GOARCH=amd64 go build -v -ldflags '-d -s -w' -o bin/nats-amd64
```

Test binary (copied from [su-exec](https://github.com/ncopa/su-exec))

```
docker run -it -v ${PWD}/bin/nats-amd64:/sbin/nats:ro alpine:latest nats
```

## TODO
* integrate tests when this project matures, like [ampq](https://github.com/streadway/amqp/blob/master/spec/gen.go)
