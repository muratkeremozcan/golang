package main

import (
  "fmt"
  "chat/util"
)

type guestConnection struct {
  ip string
  userName string
  isAdmin bool
}

func (g guestConnection)isAllowed()bool {
  return ! isIPBlocked(g.ip) && g.userName != "Darth Vader"
}

func main() {
  ip: = util.GetGuestIP()
  userName: = "Obi-Wan"

  gConn: =  & guestConnection {ip:ip, userName:userName}
  fmt.Println("Before auth", gConn)
  authorizeAdmin(gConn)
  fmt.Println("After auth", gConn)
}

func authorizeAdmin(c * guestConnection) {
  if c.isAllowed() && c.ip == "192.168.0.10" {
    c.isAdmin = true
  }
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
