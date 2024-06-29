package service

import (
	"HeroServer/proto"
	"HeroServer/service/item"
	"HeroServer/service/login"
	"HeroServer/service/lotterychest"
	"HeroServer/service/mail"
	"fmt"
	"time"
)

//type Service struct {
//	Key  uint32
//	Conn net.Conn
//}

//func NewService(key uint32, conn net.Conn) *Service {
//	service := &Service{Key: key, Conn: conn}
//	return service
//}

func HandleProto(msgId int, buf *proto.Buffer, pt *proto.Proto, PM *PlayerManager) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("HandleProto panic:", err)
			pt.MsgError("数据异常")
			time.Sleep(3 * time.Second)
			if pt.Conn != nil {
				err := pt.Conn.Close()
				if err != nil {
					return
				}
			}
		}
	}()
	switch msgId {
	case 11001:
		result := login.HandleLogin(buf, pt)
		if result != false {
			PM.OnLinePlayer[uint(pt.Rid)] = pt
		}
		break
	case 11102:
		pt.HeartBeat()
		break
	case 11101:
		pt.SendSystemTime()
		break
	case 11002:
		login.HandleStartRecord(pt)
		break
	case 41001:
		pt.RoleData.CurrencyUpdate(5, 6480)
		pt.MsgCurrencyUpdate(5, uint(pt.RoleData.CurrencyMap[5].Num))
		//pt.MsgRechargeOrder()
		//pt.MsgRechargePaid()
		pt.MsgRechargeOk()
		//pt.MPtRechargeMailAdd()
	case 18004:
		id := buf.ReadByte()
		fmt.Println("钻石id", id)
		//num := []int{60, 320, 750, 1500, 3950, 8450}

		num := map[int][]int{
			1: []int{6, 60},
			2: []int{30, 320},
			3: []int{68, 750},
			4: []int{128, 1500},
			5: []int{328, 3950},
			6: []int{648, 8450},
		}
		if pt.RoleData.CurrencyMap[5].Num >= num[int(id)][0] {
			pt.RoleData.CurrencyUpdate(1, num[int(id)][1])
			pt.RoleData.CurrencyUpdate(5, -num[int(id)][0])
			pt.MsgCurrencyUpdate(1, uint(pt.RoleData.CurrencyMap[1].Num))
			pt.MsgCurrencyUpdate(5, uint(pt.RoleData.CurrencyMap[5].Num))
			pt.MsgCurrencyBuyDiamond(id)
		} else {
			pt.MsgError("粉钻不足")
		}
		break
	case 25002:
		mail.HandleGetAllMail(pt)
		break
	case 25003:
		mail.HandleGetAttachMail(buf, pt)
		break
	case 17301:
		lotterychest.HandleOpen(buf, pt)
		break
	case 17007:
		itemkey := buf.ReadInt()
		indexlength := buf.ReadInt()
		index := []byte{}
		if indexlength != 0 {
			for i := 0; i < indexlength; i++ {
				index = append(index, buf.ReadByte())
			}
		}
		cost := [][]int{}
		costlength := buf.ReadInt()
		if costlength != 0 {
			for i := 0; i < costlength; i++ {
				cost = append(cost, []int{buf.ReadInt(), int(buf.ReadUint16())})
			}
		}
		item.WeaponRecast(itemkey, index, pt)
		fmt.Println(itemkey, index, cost)
		break
	case 17008:
		itemkey := buf.ReadInt()
		var isReplace bool
		if buf.ReadUint16() == 1 {
			isReplace = true
		} else {
			isReplace = false
		}
		item.WeaponAttrReplace(itemkey, isReplace, pt)
		break
	}
}
