package main

import (
  "fmt"
  "time"
)

type guestConnection struct {
  ip       string
  userName string
  isAdmin  bool
}

type visitorConnection struct {
  ip       string
  connHour int
}
type notifier interface {
  notify()
}


// one implementation for guestConnection
func (g guestConnection)notify() {
  fmt.Println("Guest connection from user name:", g.userName)
}

// and a different implementation for visitorConnection
func (v visitorConnection)notify() {
  fmt.Println("Visitor connected at:", v.connHour)
}

func main() {
  notifiers: = getAllConnections()
  for _, c: = range notifiers {
    c.notify()
  }

}

func getAllConnections()[]notifier {
  gConn: =  & guestConnection {ip:"192.168.0.10", userName:"Darth Vader"}
  vConn: =  & visitorConnection {ip:"192.168.0.11", connHour:time.Now().Hour()}

  return []notifier {gConn, vConn}
}

