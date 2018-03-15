package main

import (
  "fmt"
  "chat/util"
)

type guestConnection struct {
  ip string
  userName string
}

func (g guestConnection)isAllowed()bool {
  return ! isIPBlocked(g.ip) && g.userName != "Darth Vader"
}

func main() {
  ip: = util.GetGuestIP()
  userName: = "Kerry"

  gConn: = guestConnection {ip:ip, userName:userName}
  isAllowedStatus: = gConn.isAllowed(); 
  fmt.Println(isAllowedStatus)
}

func isIPBlocked(ip string)bool {
  blockedIPs: = getBlockedIPs()
  for _, blockedIP: = range blockedIPs {
    if blockedIP == ip {
      return true
    }
  }
  return false
}

func getBlockedIPs()[]string {
  return []string {"192.168.0.17", "192.168.0.18", "192.168.0.19", "192.168.0.20"}
}
