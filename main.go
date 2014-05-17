package main

import (
  "code.google.com/p/go.crypto/ssh"
  "github.com/dopuskh3/godssh/dssh"
  "log"
  "os"
)

func main() {
  if len(os.Args) != 3 {
    log.Fatalln("usage: go-ssh <group> <cmd>")
  }

  err := dssh.LoadConfigFromFile("config.yml")
  if err != nil {
    panic(err)
  }
  clientConfig, err := dssh.LoadClientConfig(os.Args[1])
  if err != nil {
    panic(err)
  }
  group, _ := dssh.GetGroup(os.Args[1])
  client, err := ssh.Dial("tcp", group.Hosts[0], clientConfig)
  if err != nil {
    panic("Failed to dial: " + err.Error())
  }
  session, err := client.NewSession()
  if err != nil {
    panic("Failed to create session: " + err.Error())
  }
  defer session.Close()

  session.Stdout = os.Stdout
  session.Stderr = os.Stderr

  if err := session.Run(os.Args[2]); err != nil {
    panic("Failed to run: " + err.Error())
  }
}
