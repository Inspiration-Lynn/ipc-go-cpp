package isiface

// 定义ipc stub的抽象接口
type IIpcConnection interface {
	// 启动ipc stub
	Start()
	// 运行ipc stub
	StubRun()
	// 停止ipc stub
	Stop()
	// 向dds发送数据
	SendDataToDDS([]byte)
	// // 设置当前连接绑定的socket conn
	// setConnection(net.Conn) 
	// // 获取当前连接绑定的socket conn
	// getConnection() net.Conn
}