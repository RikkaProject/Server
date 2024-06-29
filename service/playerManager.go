package service

import (
	"HeroServer/proto"
	"net"
)

type PlayerManager struct {
	OnLinePlayer map[uint]*proto.Proto
	TcpListener  net.Listener
}
