package main

import (
  "chat/util"
  "fmt"
)

func main() {
  hostIP: = util.GetHostIP()
  fmt.Println(isIPBlocked(hostIP))
}

func isIPBlocked(ip string)bool {
  blockedIPs: = getBlockedIPs()
  for _, blockedIP: = range blockedIPs {
    if  blockedIP == ip {
      return true
    }
  }
  return false
}

func getBlockedIPs()[]string {
  return []string {"192.168.0.17", "192.168.0.18", "192.168.0.19", "192.168.0.20"}
}
