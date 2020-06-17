package main

import (
  "log"
  "os"
  "github.com/armon/go-socks5"
)

func main() {
  conf := &socks5.Config{Logger: log.New(os.Stdout, "", log.LstdFlags)}
  server, err := socks5.New(conf)
  if err != nil {
    panic(err)
  }

  // Create SOCKS5 proxy on localhost port 1080
  log.Println("Starting socks5 server on :1080")
  if err := server.ListenAndServe("tcp", "0.0.0.0:1080"); err != nil {
    panic(err)
  }

}
