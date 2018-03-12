package main

import (
  "chat/util"
)

func main() {
  var blockedIPs []string  

  blockedIPs = append(blockedIPs, "192.168.0.16")
  blockedIPs = append(blockedIPs, "192.168.0.17")
  blockedIPs = append(blockedIPs, "192.168.0.18")
  blockedIPs = append(blockedIPs, "192.168.0.19")
  blockedIPs = append(blockedIPs, "192.168.0.20")

  addToBlockedList(blockedIPs)
}

func addToBlockedList(ips []string) {
  util.SaveBlockedIPs(ips)
}
