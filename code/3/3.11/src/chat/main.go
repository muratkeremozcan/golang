package main

import (
  "chat/util"
)

func main() {
  blockedIPs: = []string {"192.168.0.19", "192.168.0.20"}
  addToBlockedList(blockedIPs)
}

func addToBlockedList(ips []string) {
  util.SaveBlockedIPs(ips)
}
