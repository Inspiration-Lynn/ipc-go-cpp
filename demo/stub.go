package main

import (
	"ipc-go-cpp/ipcStub"
)

func main() {
	// 创建一个ipc stub句柄
	ipcStub := ipcStub.NewIpcStub()

	// 运行stub
	ipcStub.StubRun()

	// 发送信息
	// ipcStub.SendDataToDDS([]byte("position: [123, 456]"))
}