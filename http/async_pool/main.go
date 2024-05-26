package main

import (
  "net"
  "fmt"
  "os"
)

func main() {

  if len(os.Args) != 2 {
    fmt.Println("please provide a port no") 
    os.Exit(1)
  }

  ls, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Args[1]))
  c := make(chan net.Conn, 256)

  if err != nil {
    panic(err)
  }

  go newConnectionHandler(c)
  go newConnectionHandler(c)
  go newConnectionHandler(c)
  go newConnectionHandler(c)
  go newConnectionHandler(c)


  for {
    conn, err := ls.Accept()

    if err != nil {
      panic(err)
    }

    c <- conn    
  }

}

func newConnectionHandler(cc <-chan net.Conn) {
  
  for {
    select {
    case c := <-cc:
      handleConnection(c)
    }
  }
}

func handleConnection(conn net.Conn) {
    response := "{\"message\": \"Hello, World\"}"
    fmt.Fprintf(conn, "HTTP/1.1 200\r\n")
    fmt.Fprintf(conn, "Content-Type: application/json\r\n")
    fmt.Fprintf(conn, fmt.Sprintf("Content-Length: %d\r\n\r\n", len(response)))
    fmt.Fprintf(conn, response)
    conn.Close()
}
