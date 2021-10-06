package main

import (
  "net"
  "fmt"
)

var dict = map[string]string{
  "0" : "zero",
  "1" : "one",
  "2" : "two",
  "3" : "three",
  "4" : "four",
  "5" : "five",

}

func HandleConn(conn net.Conn){
  defer conn.Close()
  for {
    input := make([]byte, (1024*4))
    n, err := conn.Read(input)
    if n == 0 || err != nil {
            fmt.Println("Read error:", err)
            break
    }
    source := string(input[0:n])
    target, ok := dict[source]
    if ok == false{
      target = "undefined"
    }
    fmt.Println(source, "-", target)
    conn.Write([]byte(target))
  }
}


func main(){
  listener, err := net.Listen("tcp", ":4545")
  if err != nil {
        fmt.Println(err)
        return
    }
    defer listener.Close()
    fmt.Println("Server is listening...")
    for {
        conn, err := listener.Accept()
        if err != nil {
              fmt.Println(err)
              conn.Close()
              continue
        }
        go HandleConn(conn)
    }
}
