package main

import (
  "net"
  "fmt"
)

func main() {
  
  ls, err := net.Listen("tcp","localhost:8080")

  if err != nil {
    panic(err)
  }

  for {
    conn, err := ls.Accept()

    if err != nil {
      panic(err)
    }

    response := "{\"message\": \"Hello, World\"}"
    fmt.Fprintf(conn, "HTTP/1.1 200\r\n")
    fmt.Fprintf(conn, "Content-Type: application/json\r\n")
    fmt.Fprintf(conn, fmt.Sprintf("Content-Length: %d\r\n\r\n", len(response)))
    fmt.Fprintf(conn, response)
    conn.Close()
  }

  
}
