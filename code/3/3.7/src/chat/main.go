package main

import (
  "chat/util"
)

func main() {
  var blockedIPs [4]string
  blockedIPs [0] = "192.168.0.16"
  blockedIPs [1] = "192.168.0.17"
  blockedIPs [2] = "192.168.0.18"
  blockedIPs [3] = "192.168.0.19"
  
  addToBlockedList(blockedIPs)
}

func addToBlockedList(ips [4]string ) {
  util.SaveBlockedIPs(ips)
}
