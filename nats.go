package main

import (
    "github.com/nats-io/go-nats"
    "time"
    "os"
    "fmt"
  "gopkg.in/urfave/cli.v2" // imports as package "cli"
)

func main() {
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
          fmt.Println("Connected to server: ", c.String("server"))

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

          fmt.Println("Pushed to channel: ", channel)
          fmt.Println("Message: ", message)
          return nil
        },
      },
      {
        Name:    "sub",
        Usage:   "Subscribe to NATS channel",
        ArgsUsage: "channel",
        Action:  func(c *cli.Context) error {
          nc, _ := nats.Connect(c.String("server"))
          fmt.Println("Connected to server: ", c.String("server"))

          channel := "foo"
          if c.NArg() > 0 {
            channel = c.Args().First() // Get(0)
          }

          // Simple Async Subscriber
          nc.Subscribe(channel, func(m *nats.Msg) {
              fmt.Printf("Received a message: %s\n", string(m.Data))
          })
          fmt.Printf("Listening for messages on: %s\n", channel)
          select {} // Block forever
        },
      },
    },

  }

  app.Run(os.Args)
}
