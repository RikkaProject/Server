package main

import (
	"HeroServer/db"
	"HeroServer/gateway"
	"HeroServer/proto"
	"HeroServer/service"
	"embed"
	"fmt"
	"net"
	"strconv"
	"sync"
)

var PlayerManager service.PlayerManager

var version = "V0.1.0"

var serverWaitGroup sync.WaitGroup

//go:embed gateway/html/index.html
var html string

//go:embed gateway/html
var staticAssets embed.FS

func handleConnection(conn net.Conn) {
	defer conn.Close()
	pt := &proto.Proto{Key: 441389361, Conn: conn}
	for {
		// 读取客户端发来的数据
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			if uint(pt.Rid) > 0 {
				delete(PlayerManager.OnLinePlayer, uint(pt.Rid))
			}
			return
		}

		buf := proto.NewBufferFrom(buffer[:n])

		buflen := buf.ReadInt()
		msgId := buf.ReadInt() //消息id

		fmt.Println("BufLen:" + strconv.Itoa(buflen) + " MSG ID:" + strconv.Itoa(msgId) + " TotalLen:" + strconv.Itoa(n))
		fmt.Println(buf.Byte)

		//buf.ResetOffset()
		service.HandleProto(msgId, buf, pt, &PlayerManager)
	}
}

func main() {
	serverWaitGroup.Add(2)
	db.InitDB()

	PlayerManager = service.PlayerManager{
		OnLinePlayer: make(map[uint]*proto.Proto),
	}

	go gateway.HandleHttp(&PlayerManager, version, html, staticAssets)
	go startServer()

	serverWaitGroup.Wait()
}

func startServer() {
	var err error
	PlayerManager.TcpListener, err = net.Listen("tcp", ":6001")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer PlayerManager.TcpListener.Close()

	fmt.Println("[RikkaServer] Listening on :6001")
	fmt.Println("[RikkaServer] 我的勇者, 启动！！！")

	for {
		conn, err := PlayerManager.TcpListener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		fmt.Println("Client connected:", conn.RemoteAddr())
		go handleConnection(conn)
	}
}
