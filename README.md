# nats-cli [![Build Status](https://travis-ci.org/shadiakiki1986/nats-cli.svg?branch=master)](https://travis-ci.org/shadiakiki1986/nats-cli)
NATS client CLI

Equivalent to the [ruby-nats](https://github.com/nats-io/ruby-nats) example scripts

## Usage


## Installation

Download the binary from one of the [releases](https://github.com/shadiakiki1986/nats-cli/releases)
(in example below, `amd64` is the output of `dpkg --print-architecture`)

```bash
wget https://github.com/shadiakiki1986/nats-cli/releases/download/0.0.3/nats-amd64 -O /sbin/nats
chmod +x /sbin/nats
```
Now you can run it with:
```bash
nats ...
```

## Usage
1. Publish to channel "foo" the message "help me!"
```bash
bin/nats pub foo "help me!"
```

2. Do the same on a different server
```bash
bin/nats --server nats://someserver:4222 pub foo "help me!"
```

3. Subscribe to channel "foo" and just display the messages received in the console
```bash
bin/nats sub foo
```

4. Subscribe to channel "foo" and trigger a command upon receipt of the generated token

```bash
shadi@ffa24:~/Development/nats-cli$ go run nats.go sub --cmd 'echo "hey"' foo
2017/02/03 10:33:19 Start
2017/02/03 10:33:19 Connected to server:  nats://localhost:4222
2017/02/03 10:33:19 Listening for messages on: foo
2017/02/03 10:33:19 Message should match with token: 022d59bdf309dc22
2017/02/03 10:33:25 Received a message: 022d59bdf309dc22
2017/02/03 10:33:25 Messag matches with token .. triggering command: 'echo "hey"'
2017/02/03 10:33:25 >>>
hey

<<<
2017/02/03 10:33:25 Listening for messages on: foo
```

In the example above, the token is sent using

```bash
shadi@ffa24:~/Development/nats-cli$ go run nats.go pub foo 022d59bdf309dc22
2017/02/03 10:33:25 Start
2017/02/03 10:33:25 Connected to server:  nats://localhost:4222
2017/02/03 10:33:25 Pushed to channel:  foo
2017/02/03 10:33:25 Message:  022d59bdf309dc22
```

## Releasing
1. Pre-requisites
```bash
sudo apt-get install golang
export GOPATH=${PWD}

# https://github.com/nats-io/go-nats
go get github.com/nats-io/go-nats

# https://github.com/urfave/cli#using-the-v2-branch
go get gopkg.in/urfave/cli.v2
```

2. Edit version number in nats.go

3. Build binary (copied from [gosu](https://github.com/tianon/gosu/blob/master/Dockerfile))

```bash
CGO_ENABLED=0 GOARCH=amd64 go build -v -ldflags '-d -s -w' -o bin/nats-amd64
```

4. Test binary (copied from [su-exec](https://github.com/ncopa/su-exec))

```
docker run -it -v ${PWD}/bin/nats-amd64:/sbin/nats:ro alpine:latest nats
```

5. Add git tag and push

6. Go to github releases, edit the tag, and upload the newly built binary

## TODO
* integrate tests when this project matures, like [ampq](https://github.com/streadway/amqp/blob/master/spec/gen.go)
