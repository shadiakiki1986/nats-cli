# nats-cli
NATS client CLI

Equivalent to the ruby-nats example scripts

## Usage

Download one of the releases

Run

```
nats-cli pub foo "help me!"
```

## Development
```bash
sudo apt-get install golang
export GOPATH=${PWD}
go get github.com/nats-io/go-nats
go get gopkg.in/urfave/cli.v2
```

## TODO
binary wont run .. check https://github.com/tianon/gosu/blob/master/Dockerfile
