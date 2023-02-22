package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const socketPath = "/home/lynn/iCube/ipc/ipcSocket1.sock"
 
func main() {
    // Create a Unix domain socket and listen for incoming connections.
    listener, err := net.Listen("unix", socketPath)
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    for {
        // Accept an incoming connection.
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Start stub successfully...")

        // Handle the connection in a separate goroutine.
        go handleConn(conn)
    }
}

func handleConn(conn net.Conn) {
    defer conn.Close()
    

    for {
        // Create a buffer for incoming data.
        buf := make([]byte, 4096)   

        // Read data from the connection.
        cnt, err := conn.Read(buf)
        if err == io.EOF {
			break
		}
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("recv data: %s, cnt=%d\n", buf[:cnt-1], cnt)

        // Echo the data back to the connection.
        _, err = conn.Write(buf[:cnt])
        if err != nil {
            log.Fatal(err)
        }
    }
}