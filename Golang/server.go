package main

import (
  "net"
  "fmt"
)

func main(){
  message := "Hello I am server"
  listener, err := net.Listen("tcp", ":4545")

  if err != nil{
    fmt.Println(err)
    return
  }
  defer listener.Close()
  fmt.Println("Server is listening...")
  for {
    conn, err := listener.Accept()
    if err != nil{
      fmt.Println(err)
      return
    }
    conn.Write([]byte(message))
    conn.Close()
  }
}
