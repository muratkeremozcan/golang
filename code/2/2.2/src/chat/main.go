package main

import (
  "chat/util"
  "os"
)

func main() {
  listenPort := "8080"
  args := os.Args
  if len(args) > 1 {
    hostIP := args[1]
    util.RunGuest(hostIP)
  } else {
    listenIP := util.GetLocalNetworkIP()
    util.RunHost(listenIP + ":" + listenPort)
  }
}
