package main

import (
  "fmt"
)

func main() {
  list := getBlockedIPs()
  for i := range list {
    fmt.Println(list[i])
  }
}

func getBlockedIPs() []string {
  return []string{"192.168.0.17", "192.168.0.18", "192.168.0.19", "192.168.0.20"}
}
