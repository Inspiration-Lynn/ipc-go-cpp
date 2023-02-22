package ipcStub

import (
	"fmt"
	"io"
	"ipc-go-cpp/isiface"
	"log"
	"net"
	"os"
)

// ipc stub接口实现
type IpcConnection struct {
	SocketPath string
	Connection net.Conn
}

// const sockP = "/home/lynn/iCube/ipc/ipcSocket1.sock"

// 初始化ipcstub模块的方法
func NewIpcStub() isiface.IIpcConnection{
	return &IpcConnection {
		SocketPath: "/home/lynn/go/src/ipc-go-cpp/demo/ipcSocket1.sock",
		Connection: nil,
	}
}

// 启动ipc stub
func (ic *IpcConnection) Start() {
	go func() {
		os.Remove(ic.SocketPath)
		listener, err := net.Listen("unix", ic.SocketPath)
		if err != nil {
			log.Fatal(err)
		}
		defer listener.Close()
		fmt.Println("[Start ipc stub successfully]")
	
		// 阻塞等待proxy侧的连接
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			ic.Connection = conn
			fmt.Println("[Connection is established]")
			fmt.Println("local addr: ", conn.LocalAddr().String())
	
			go func(c net.Conn) {
				for {
					buf := make([]byte, 1024)
					_, err := c.Read(buf)
					// 连接断开，退出此goroutine
					if err != nil && err != io.EOF {
						c.Close()
						break
					}
					if err == io.EOF {
						fmt.Println("recv EOF")
						break
					}
					fmt.Println("recv: ", string(buf))
					c.Write([]byte("[ok]"))
				}
				c.Close()
			}(conn)
		}
	}()
}

// 运行ipc stub
func (ic *IpcConnection) StubRun() {
	// Start非阻塞
	ic.Start()

	// 阻塞状态
	select{}
}

// 停止ipc stub
func (ic *IpcConnection) Stop() {
	fmt.Println("server stop...")
	// 资源回收
}

// // 设置当前连接绑定的socket conn
// func (ic *IpcConnection) setConnection(conn net.Conn) {
// 	ic.Connection = conn
// }

// // 获取当前连接绑定的socket conn
// func (ic *IpcConnection) getConnection() (net.Conn) {
// 	return ic.Connection
// }

func (ic *IpcConnection) SendDataToDDS(data []byte) {
	// conn := ic.getConnection()
	if _, err := ic.Connection.Write(data); err != nil {
		log.Print(err)
		return
	// }
	// var buf = make([]byte, 1024)
	// if _, err := conn.Read(buf); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("client recv: ", string(buf))
	}
}
