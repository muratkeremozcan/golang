package main

import "fmt"

type guestConnection struct {
  ip string
  userName string
  isAdmin bool
}

func (g guestConnection) notify() {
  fmt.Println("Guest connection from user name:", g.userName)
}

func main() {
  guestConns := getAllConnections()
  for _, c := range guestConns {
    c.notify()
  }
}

func getAllConnections() []*guestConnection  {
  gConn1 := &guestConnection{ip: "192.168.0.10", userName: "Darth Vader"}
  gConn2 := &guestConnection{ip: "192.168.0.11", userName: "Obi-Wan"}
  
  return []*guestConnection{gConn1, gConn2}
}
