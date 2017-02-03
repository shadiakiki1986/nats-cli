package main

import (
    "github.com/nats-io/go-nats"
    "time"
    "os"
    "os/exec"
    "fmt"
    "gopkg.in/urfave/cli.v2" // imports as package "cli"
    "crypto/rand"
    "log"
)

// http://stackoverflow.com/a/25431798/4126114
func randToken() string {
    b := make([]byte, 8)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}

func main() {
  // Go: Global variables
  // http://stackoverflow.com/a/25096729/4126114
  var token = ""

  // golang log example
  // https://golang.org/pkg/log/#example_Logger
  // using golang global logger
  // http://stackoverflow.com/a/18361927/4126114
  log.Printf("Start")

  app := &cli.App{
    Flags: []cli.Flag {
      &cli.StringFlag{
        Name:        "server",
        Value:       nats.DefaultURL, // "nats://localhost:4222",
        Usage:       "NATS server URI",
      },
    },
    Commands: []*cli.Command{
      {
        Name:    "pub",
        Usage:   "Publish to NATS channel",
        ArgsUsage: "channel message",
        Action:  func(c *cli.Context) error {
          nc, _ := nats.Connect(c.String("server"))
          log.Println("Connected to server: ", c.String("server"))

          channel := "foo"
          if c.NArg() > 0 {
            channel = c.Args().First() // Get(0)
          }
          message := "help me!"
          if c.NArg() > 1 {
            message = c.Args().Get(1)
          }

          // Make a request
          nc.Request(channel, []byte(message), 10*time.Millisecond)

          log.Println("Pushed to channel: ", channel)
          log.Println("Message: ", message)
          return nil
        },
      },
      {
        Name:    "sub",
        Usage:   "Subscribe to NATS channel",
        ArgsUsage: "channel",
        Flags: []cli.Flag {
          &cli.StringFlag{
            Name:        "cmd",
            Value:       "",
            Usage:       "Command to run upon receipt of displayed token on channel",
          },
          &cli.StringFlag{
            Name:        "token",
            Value:       "",
            Usage:       "Token to listen for in order to run the command",
          },
        },
        Action:  func(c *cli.Context) error {
          nc, _ := nats.Connect(c.String("server"))
          log.Println("Connected to server: ", c.String("server"))

          channel := "foo"
          if c.NArg() > 0 {
            channel = c.Args().First() // Get(0)
          }

          // check if cmd is passed
          // http://stackoverflow.com/a/25431798/4126114
          if c.String("cmd")!="" {
            // check if token is provided
            if c.String("token")=="" {
              // generate a random token
              token = randToken()
            } else {
              token = c.String("token")
            }
          }

          // Simple Async Subscriber
          nc.Subscribe(channel, func(m *nats.Msg) {
              log.Printf("Received a message: %s\n", string(m.Data))
              if c.String("cmd")!="" {
                if string(m.Data) == token {
                  // How to execute system command in Golang with unknown arguments
                  // http://stackoverflow.com/a/20438245/4126114
                  log.Printf("Message matches with token .. triggering command: '%s'\n", c.String("cmd"));
                  out, err := exec.Command("sh","-c",c.String("cmd")).Output()
                  if err != nil {
                      log.Println("error occured")
                      log.Printf("%s", err)
                  }
                  log.Printf(">>>\n%s\n<<<", out)
                }
              }

             log.Printf("Listening for messages on: %s\n", channel)
          })
          log.Printf("Listening for messages on: %s\n", channel)
          if c.String("cmd")!="" {
            log.Printf("Message should match with token: %s\n",token);
          }

          select {} // Block forever
        },
      },
    },

  }

  app.Name = "nats"
  app.Version = "0.0.4.1"
  app.Run(os.Args)
}
