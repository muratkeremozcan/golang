package main

import (
  "fmt"
  "pinger/util"
)

func main() {
  isAlive: = false
  for ! isAlive {
    if util.PingChatServer() {
      isAlive = true
    }
  }
  
  fmt.Println("Ping is done")
}
