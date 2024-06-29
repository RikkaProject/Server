package proto

import (
	"HeroServer/db"
	"fmt"
	"time"
)

func (p *Proto) Login(role *db.Role) {
	sbuf := NewBuffer(1024, p)
	sbuf.PutUint16(11005)
	sbuf.PutInt(1)                    //数量?
	sbuf.PutInt(role.RecordId)        //recordId
	sbuf.PutByte(byte(role.RecordId)) //recordIndex
	sbuf.PutInt(role.RoleId)          //roleId
	sbuf.PutInt(role.Level)           //等级
	sbuf.PutInt(14167)                //MainWeaponCfgId
	sbuf.PutByte(0)                   //不知道
	{                                 //时装 costumeIDList
		sbuf.PutInt(2)
		sbuf.PutInt(26992)
		sbuf.PutInt(26993)
	}
	sbuf.PutInt(0)
	sbuf.Finish(false)
	p.Conn.Write(sbuf.Byte)

	sbuf2 := NewBuffer(1024, p)
	sbuf2.PutUint16(11015)
	sbuf2.PutInt(0)
	sbuf2.Finish(false)
	p.Conn.Write(sbuf2.Byte)

	buf3 := NewBuffer(17, p)
	buf3.PutUint16(11101)
	buf3.PutLong(int(time.Now().UnixNano() / int64(time.Millisecond)))
	buf3.PutByte(8)
	buf3.Finish(false)
	p.Conn.Write(buf3.Byte)

	buf4 := NewBuffer(17, p)
	buf4.PutUint16(11020)
	buf4.PutByte(1)
	buf4.Finish(false)
	p.Conn.Write(buf4.Byte)

	//p.Conn.Write([]byte{57, 0, 14, 47, 0, 24, 0, 0, 0, 1, 0, 2, 0, 3, 0, 6, 0, 7, 0, 8, 0, 10, 0, 11, 0, 12, 0, 13, 0, 15, 0, 16, 0, 17, 0, 19, 0, 21, 0, 23, 0, 24, 0, 25, 0, 27, 0, 28, 0, 29, 0, 30, 0, 34, 0, 35, 0, 85, 0, 249, 42, 162, 149, 217, 253, 187, 108, 2, 56, 0, 0, 0, 0, 109, 23, 81, 42, 21, 0, 0, 0, 233, 173, 148, 230, 179, 149, 229, 176, 145, 229, 165, 179, 228, 188, 138, 232, 142, 137, 233, 155, 133, 0, 4, 1, 32, 0, 0, 0, 48, 48, 56, 53, 50, 97, 49, 98, 56, 50, 98, 57, 98, 100, 53, 54, 55, 101, 97, 98, 52, 98, 55, 55, 97, 48, 97, 101, 53, 53, 100, 51, 0})
	buf5 := NewBuffer(1024, p)
	buf5.PutUint16(12046)
	buf5.PutByte(0)
	buf5.PutInt(24)
	buf5.PutUint16(1)
	buf5.PutUint16(2)
	buf5.PutUint16(3)
	buf5.PutUint16(6)
	buf5.PutUint16(7)
	buf5.PutUint16(8)
	buf5.PutUint16(10)
	buf5.PutUint16(11)
	buf5.PutUint16(12)
	buf5.PutUint16(13)
	buf5.PutUint16(15)
	buf5.PutUint16(16)
	buf5.PutUint16(17)
	buf5.PutUint16(19)
	buf5.PutUint16(21)
	buf5.PutUint16(23)
	buf5.PutUint16(24)
	buf5.PutUint16(25)
	buf5.PutUint16(27)
	buf5.PutUint16(28)
	buf5.PutUint16(29)
	buf5.PutUint16(30)
	buf5.PutUint16(34)
	buf5.PutUint16(35)
	buf5.Finish(false)
	p.Conn.Write(buf5.Byte)

	buf6 := NewBuffer(1024, p)
	buf6.PutUint16(11001)
	buf6.PutLong(int(role.PlayerKey))
	buf6.PutFloat(0)       //version
	buf6.PutInt(709957485) //seed
	buf6.PutString(role.Name)
	buf6.PutByte(1)                                    //renameFree
	buf6.PutByte(4)                                    //lastLoginRole
	buf6.PutByte(1)                                    //region
	buf6.PutString("00852a1b82b9bd567eab4b77a0ae55d3") //quickToken
	buf6.PutByte(1)                                    //isQuickLogin
	buf6.Finish(false)
	p.Conn.Write(buf6.Byte)
}

// SendRoleInfo 返回角色消息
func (p *Proto) SendRoleInfo() {
	buf := NewBuffer(1024, p)
	//buf.PutInt(2705096883)
	buf.PutUint16(12011)                        // msgId
	buf.PutLong(int(p.RoleData.Role.PlayerKey)) // playerKey
	buf.PutString(p.RoleData.Role.Name)         // name 名字
	buf.PutString(p.RoleData.Role.SelfIntro)    //SelfIntro 简介
	buf.PutInt(p.RoleData.Role.RecordId)        //recordId 记录id 1战士 2游侠 3法师 4牧师 5勇者
	buf.PutInt(p.RoleData.Role.RoleId)          //roleId 角色id 1战士 2游侠 3法师 4牧师 5勇者
	buf.PutInt(p.RoleData.Role.Level)           //lv 等级
	buf.PutInt(p.RoleData.Role.Exp)             //exp 经验
	buf.PutInt(1709174476)                      //roleCreateTime 角色创建时间
	buf.PutInt(p.RoleData.Role.Id)              //BraveId 勇者ID

	buf.Finish(true)
	fmt.Println("SendRoleInfo", buf.Byte)
	p.Conn.Write(buf.Byte)
}

// SendCurrencyNumber 应该是背包材料 第一个Int 应该为总数 然后一个Int(物品id)和一个Uint(数量)为一组
func (p *Proto) SendCurrencyNumber(items []*db.Currency) {
	buf := NewBuffer(81920, p)
	buf.PutUint16(18001) // msgId
	buf.PutInt(len(items))
	p.RoleData.CurrencyMap = map[int]*db.Currency{}
	if len(items) > 0 {
		for _, item := range items {
			p.RoleData.CurrencyMap[item.ItemId] = item
			buf.PutInt(item.ItemId)
			buf.PutInt(item.Num)
		}
	}

	buf.Finish(true)

	//fmt.Println("SendCurrencyNumber", buf.Byte)
	p.Conn.Write(buf.Byte)

}

// MsgRechargeOk 可以弹出奖励窗口可 以不只是粉钻
func (p *Proto) MsgRechargeOk() {
	buf := NewBuffer(1024, p)
	buf.PutUint16(41003)

	buf.PutInt(201)  //Recharge
	buf.PutUint16(4) //Proto.MpsReceiveInfo.ReceiveFlag
	buf.PutInt(2)    //物品列表数量
	buf.PutInt(5)    //id
	buf.PutInt(3240) //物品数量
	buf.PutInt(5)    //id
	buf.PutInt(3240) //物品数量

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

func (p *Proto) MsgRechargeOrder() {
	buf := NewBuffer(1024, p)
	buf.PutUint16(41001)
	buf.PutUint16(0)
	buf.PutInt(201)
	buf.PutString("202403072133577001967686")

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

//func (p *Proto) MsgRechargePaid() { 好像是取消支付
//	buf := NewBuffer(1024, p)
//	buf.PutUint16(41002)
//	buf.PutUint16(0)
//	buf.PutUint16(199)
//	buf.PutString("321123")
//
//	buf.Finish(true)
//	p.Conn.Write(buf.Byte)
//}

func (p *Proto) MPtRechargeMailAdd() {
	buf := NewBuffer(1024, p)
	buf.PutUint16(41103)
	buf.PutInt(1)
	buf.PutString("321123")
	buf.PutInt(199)
	buf.PutInt(int(time.Now().Unix()))

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

func (p *Proto) MsgCurrencyBuyDiamond(id byte) {
	buf := NewBuffer(1024, p)
	buf.PutUint16(18004)
	buf.PutUint16(0)
	buf.PutByte(id)

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

func (p *Proto) MsgCurrencyUpdate(itemId int, num uint) {
	buf := NewBuffer(1024, p)
	buf.PutUint16(18002)
	buf.PutInt(1)
	buf.PutInt(itemId)
	buf.PutInt(int(num))

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

// MsgHeadInfo 头像信息
func (p *Proto) MsgHeadInfo() {
	buf := NewBuffer(1024, p)
	buf.PutUint16(51003) // msgId

	{ //头像
		buf.PutInt(28)
		buf.PutInt(2001)
		buf.PutInt(0)
		buf.PutInt(2002)
		buf.PutInt(0)
		buf.PutInt(2003)
		buf.PutInt(0)
		buf.PutInt(2004)
		buf.PutInt(0)
		buf.PutInt(2005)
		buf.PutInt(0)
		buf.PutInt(2006)
		buf.PutInt(0)
		buf.PutInt(2007)
		buf.PutInt(0)
		buf.PutInt(2008)
		buf.PutInt(0)
		buf.PutInt(2009)
		buf.PutInt(0)
		buf.PutInt(2010)
		buf.PutInt(0)
		buf.PutInt(2011)
		buf.PutInt(0)
		buf.PutInt(2012)
		buf.PutInt(0)
		buf.PutInt(2013)
		buf.PutInt(0)
		buf.PutInt(2014)
		buf.PutInt(0)
		buf.PutInt(2015)
		buf.PutInt(0)
		buf.PutInt(2016)
		buf.PutInt(0)
		buf.PutInt(2017)
		buf.PutInt(0)
		buf.PutInt(2018)
		buf.PutInt(0)
		buf.PutInt(2019)
		buf.PutInt(0)
		buf.PutInt(2020)
		buf.PutInt(0)
		buf.PutInt(2021)
		buf.PutInt(0)
		buf.PutInt(2022)
		buf.PutInt(0)
		buf.PutInt(2023)
		buf.PutInt(0)
		buf.PutInt(2024)
		buf.PutInt(0)
		buf.PutInt(2025)
		buf.PutInt(0)
		buf.PutInt(2026)
		buf.PutInt(0)
		buf.PutInt(2035)
		buf.PutInt(0)
		buf.PutInt(2036)
		buf.PutInt(0)
	}
	{ //头像框
		buf.PutInt(17)
		buf.PutInt(1001)
		buf.PutInt(0)
		buf.PutInt(1002)
		buf.PutInt(0)
		buf.PutInt(1003)
		buf.PutInt(0)
		buf.PutInt(1004)
		buf.PutInt(0)
		buf.PutInt(1005)
		buf.PutInt(0)
		buf.PutInt(1006)
		buf.PutInt(0)
		buf.PutInt(1007)
		buf.PutInt(0)
		buf.PutInt(1008)
		buf.PutInt(0)
		buf.PutInt(1009)
		buf.PutInt(0)
		buf.PutInt(1010)
		buf.PutInt(0)
		buf.PutInt(1026)
		buf.PutInt(0)
		buf.PutInt(1033)
		buf.PutInt(0)
		buf.PutInt(1049)
		buf.PutInt(0)
		buf.PutInt(1146)
		buf.PutInt(0)
		buf.PutInt(1155)
		buf.PutInt(0)
		buf.PutInt(1179)
		buf.PutInt(0)
		buf.PutInt(1226)
		buf.PutInt(0)
	}
	{ //聊天框
		buf.PutInt(0)
	}
	{ //荣耀徽章
		buf.PutInt(1)
		buf.PutInt(502000)
		buf.PutInt(0)
	}

	buf.Finish(true)

	fmt.Println("MsgHeadInfo", buf.Byte)
	p.Conn.Write(buf.Byte)
}

// MPtAdventureLevelInfo 勇者等级信息
func (p *Proto) MPtAdventureLevelInfo() {
	buf := NewBuffer(1024, p)
	buf.PutUint16(22020) // msgId

	buf.PutInt(26)   //勇者等级
	buf.PutInt(3900) //勇气值
	buf.PutInt(26)   //已经领取等级奖励id数量
	buf.PutByte(26)  //已领取的等级奖励id
	buf.PutByte(25)
	buf.PutByte(24)
	buf.PutByte(23)
	buf.PutByte(22)
	buf.PutByte(15)
	buf.PutByte(5)
	buf.PutByte(6)
	buf.PutByte(7)
	buf.PutByte(8)
	buf.PutByte(9)
	buf.PutByte(10)
	buf.PutByte(11)
	buf.PutByte(12)
	buf.PutByte(13)
	buf.PutByte(14)
	buf.PutByte(16)
	buf.PutByte(17)
	buf.PutByte(18)
	buf.PutByte(19)
	buf.PutByte(20)
	buf.PutByte(21)
	buf.PutByte(4)
	buf.PutByte(3)
	buf.PutByte(2)
	buf.PutByte(1)

	buf.Finish(true)

	fmt.Println("MPtAdventureLevelInfo", buf.Byte)
	p.Conn.Write(buf.Byte)
}

// MPtAdventureWeeklyRewardInfo 勇者等级每周奖励
func (p *Proto) MPtAdventureWeeklyRewardInfo() {
	buf := NewBuffer(1024, p)
	buf.PutUint16(22022) // msgId
	buf.PutByte(0)
	buf.Finish(true)

	fmt.Println("MPtAdventureWeeklyRewardInfo", buf.Byte)
	p.Conn.Write(buf.Byte)
}

func (p *Proto) MsgError(err string) {
	buf := NewBuffer(1024, p)
	buf.PutUint16(998)
	buf.PutString(err)
	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

// SendItemAll 返回所有物品 暂未解出格式返回个死数据
func (p *Proto) SendItemAll(items []*db.HeroItem) {
	buf := NewBuffer(102400, p)

	/*
		{   武器数据解析 幽骸
			buf.PutInt(749) itemKey
			buf.PutInt(14167) itemId
			buf.PutByte(101) type
			buf.PutByte(0) _isStoraged
			buf.PutByte(5) quality
			buf.PutByte(5) star
			buf.PutByte(60) lv

			buf.PutInt(35821) exp
			buf.PutUint16(0) role

			buf.PutInt(3) attrListlength
			buf.PutByte(51) 属性id
			buf.PutUint16(1000) 数值百分比 1000是100%
			buf.PutByte(36)
			buf.PutUint16(1000)
			buf.PutByte(8)
			buf.PutUint16(1000)

			buf.PutByte(0) rindex
			buf.PutUint16(0) rRole

			buf.PutInt(0) rAttrListlength
			buf.PutByte(0) isNew
			buf.PutInt(0) bigRuneslength
			buf.PutInt(0) smallRuneslength
			buf.PutInt(0) coreRuneAffixlength
			buf.PutInt(0) crysRuneAffixlength
			buf.PutUint16(0) season

			buf.PutByte(1) 确定是IsProtected

			buf.PutUint16(0) RunePunchFailedTime
			buf.PutByte(0) 是否觉醒

			buf.PutInt(0) 觉醒属性数量 SpecAttrList
			buf.PutInt(0) StampAttrListlength

			buf.PutInt(3) 宝石孔数量
			buf.PutByte(1) 宝石孔位置
			buf.PutInt(3000) 宝石id
			buf.PutByte(2)
			buf.PutInt(2925)
			buf.PutByte(3)
			buf.PutInt(2881)

			buf.PutInt(0) SpcHoleInfos？
		}*
	*/

	buf.PutUint16(17001) // msgId
	buf.PutInt(len(items))
	p.RoleData.ItemMap = map[int]*db.HeroItem{}
	if len(items) > 0 {
		for _, item := range items {
			p.RoleData.ItemMap[item.Id] = item
			if item.ItemType == 101 {
				buf.PutInt(item.Id)
				buf.PutInt(item.ItemId)
				buf.PutByte(byte(item.ItemType))
				buf.PutByte(byte(item.IsStoraged))
				buf.PutByte(byte(item.Quality))
				buf.PutByte(byte(item.Star))
				buf.PutByte(byte(item.Lv))
				buf.PutInt(item.Exp)
				buf.PutUint16(item.Role)
				if item.AttrList != "" {
					attrlist := item.GetAttrList()
					buf.PutInt(len(attrlist))
					if len(attrlist) > 0 {
						for _, e := range attrlist {
							buf.PutByte(byte(e[0]))
							buf.PutUint16(e[1])
						}
					}
				} else {
					buf.PutInt(0)
				}
				buf.PutByte(0)   //rindex
				buf.PutUint16(0) //rRole

				buf.PutInt(0)    //rAttrListlength
				buf.PutByte(0)   //isNew
				buf.PutInt(0)    //bigRuneslength
				buf.PutInt(0)    //smallRuneslength
				buf.PutInt(0)    //coreRuneAffixlength
				buf.PutInt(0)    //crysRuneAffixlength
				buf.PutUint16(0) //season

				buf.PutByte(byte(item.IsProtected)) //确定是IsProtected

				buf.PutUint16(0) //RunePunchFailedTime
				buf.PutByte(0)   //是否觉醒

				buf.PutInt(0) //觉醒属性数量 SpecAttrList
				buf.PutInt(0) //StampAttrListlength

				buf.PutInt(0) //宝石孔数量
				buf.PutInt(0) //SpcHoleInfos？
			}
		}
	}

	/*{
		buf.PutInt(3)
		buf.PutInt(13025)
		buf.PutByte(101)
		buf.PutByte(0)
		buf.PutByte(3)
		buf.PutByte(0)
		buf.PutByte(1)
		buf.PutInt(0)
		buf.PutUint16(4)
		buf.PutInt(3)
		buf.PutByte(3)
		buf.PutUint16(340)
		buf.PutByte(4)
		buf.PutUint16(612)
		buf.PutByte(51)
		buf.PutUint16(121)
		buf.PutByte(0)
		buf.PutUint16(0)
		buf.PutInt(0)
		buf.PutByte(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutUint16(0)
		buf.PutByte(0)
		buf.PutUint16(0)
		buf.PutByte(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(1)
		buf.PutInt(11004)
		buf.PutByte(101)
		buf.PutByte(0)
		buf.PutByte(1)
		buf.PutByte(0)
		buf.PutByte(1)
		buf.PutInt(0)
		buf.PutUint16(0)
		buf.PutInt(0)
		buf.PutByte(0)
		buf.PutUint16(0)
		buf.PutInt(0)
		buf.PutByte(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutUint16(0)
		buf.PutByte(0)
		buf.PutUint16(0)
		buf.PutByte(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(2)
		buf.PutInt(70101)
		buf.PutByte(109)
		buf.PutByte(0)
		buf.PutByte(2)
		buf.PutByte(0)
		buf.PutByte(0)
		buf.PutInt(0)
		buf.PutUint16(4)
		buf.PutInt(1)
		buf.PutByte(21)
		buf.PutUint16(708)
		buf.PutByte(0)
		buf.PutUint16(0)
		buf.PutInt(0)
		buf.PutByte(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutUint16(0)
		buf.PutByte(0)
		buf.PutUint16(0)
		buf.PutByte(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)

		buf.PutInt(749)
		buf.PutInt(14167)
		buf.PutByte(101)
		buf.PutByte(0)
		buf.PutByte(5)
		buf.PutByte(5)
		buf.PutByte(60)
		buf.PutInt(35821)
		buf.PutUint16(1)
		buf.PutInt(3)
		buf.PutByte(51)
		buf.PutUint16(1000)
		buf.PutByte(36)
		buf.PutUint16(1000)
		buf.PutByte(8)
		buf.PutUint16(1000)
		buf.PutByte(1)
		buf.PutUint16(1)
		buf.PutInt(0)
		buf.PutByte(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutUint16(0)
		buf.PutByte(1)
		buf.PutUint16(0)
		buf.PutByte(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)

		buf.PutInt(1894)
		buf.PutInt(14171)
		buf.PutByte(101)
		buf.PutByte(0)
		buf.PutByte(5)
		buf.PutByte(5)
		buf.PutByte(60)
		buf.PutInt(35821)
		buf.PutUint16(0)
		buf.PutInt(3)
		buf.PutByte(8)
		buf.PutUint16(975)
		buf.PutByte(36)
		buf.PutUint16(902)
		buf.PutByte(51)
		buf.PutUint16(926)
		buf.PutByte(0)
		buf.PutUint16(0)
		buf.PutInt(0)
		buf.PutByte(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutInt(0)
		buf.PutUint16(0)
		buf.PutByte(1)
		buf.PutUint16(0)
		buf.PutByte(1)
		buf.PutInt(1)
		buf.PutByte(51)
		buf.PutUint16(745)
		buf.PutInt(0)
		buf.PutInt(3)
		buf.PutByte(1)
		buf.PutInt(2826)
		buf.PutByte(2)
		buf.PutInt(2879)
		buf.PutByte(3)
		buf.PutInt(2864)
		buf.PutInt(0)
	}*/

	buf.Finish(true)

	fmt.Println("SendItemAll", buf.Byte)
	p.Conn.Write(buf.Byte)
}

// SendEquipInit 装备的武器防具
func (p *Proto) SendEquipInit() {
	buf := NewBuffer(1024, p)
	//buf.PutInt(1859565377)
	buf.PutUint16(29001) // msgId

	buf.PutByte(11)
	buf.PutInt(0)
	//buf.PutByte(11)
	//buf.PutInt(749)
	buf.PutInt(0)

	buf.Finish(true)

	fmt.Println("SendEquipInit", buf.Byte)
	p.Conn.Write(buf.Byte)
}

// SendNewBagInit 背包里的物品
func (p *Proto) SendNewBagInit(items []*db.HeroItem) {
	buf := NewBuffer(1024, p)
	buf.PutUint16(28001) // msgId

	buf.PutInt(len(items))
	if len(items) > 0 {
		for _, item := range items {
			buf.PutInt(item.Id)
		}
	}
	buf.PutInt(0)

	buf.Finish(true)

	fmt.Println("SendNewBagInit", buf.Byte)
	p.Conn.Write(buf.Byte)
}

// MPtMailGetAll 获取所有邮件
func (p *Proto) MPtMailGetAll(items []*db.Mail) {
	buf := NewBuffer(81920, p)
	buf.PutUint16(25002)   // msgId
	buf.PutInt(len(items)) //数量

	if len(items) > 0 {
		for _, item := range items {
			fmt.Println(item.MailItems)
			buf.PutInt(item.Id)         //id
			buf.PutInt(item.TemplateId) //模板id
			buf.PutInt(1)               //标题数量
			buf.PutString(item.Title)
			buf.PutInt(1) //内容数量
			buf.PutString(item.Content)
			{
				mailItems := item.GetMailItems()
				mailItemsLen := len(mailItems)
				buf.PutInt(mailItemsLen) //邮件物品数量
				if mailItemsLen > 0 {
					for _, e := range mailItems {
						buf.PutInt(e[0])        //物品id
						buf.PutInt(e[1])        //物品数量
						buf.PutByte(byte(e[2])) //是否已领取
					}
				}
			}

			buf.PutByte(byte(item.Status)) //状态 0未读 1已读 2已接收
			buf.PutInt(item.Time)          //过期时间
			buf.PutString("")              //图片
			buf.PutString(item.Sender)     //发件人
		}
	}

	buf.Finish(true)

	fmt.Println("MPtMailGetAll", buf.Byte)
	p.Conn.Write(buf.Byte)
}

func (p *Proto) MPtMailNewInfo(mail *db.Mail) {
	buf := NewBuffer(8192, p)
	buf.PutUint16(25005)        // msgId
	buf.PutInt(mail.Id)         //id
	buf.PutInt(mail.TemplateId) //模板id
	buf.PutInt(1)               //标题数量
	buf.PutString(mail.Title)
	buf.PutInt(1) //内容数量
	buf.PutString(mail.Content)
	{
		mailItems := mail.GetMailItems()
		mailItemsLen := len(mailItems)
		buf.PutInt(mailItemsLen) //邮件物品数量
		if mailItemsLen > 0 {
			for _, e := range mailItems {
				buf.PutInt(e[0])        //物品id
				buf.PutInt(e[1])        //物品数量
				buf.PutByte(byte(e[2])) //是否已领取
			}
		}
	}

	buf.PutByte(byte(mail.Status)) //状态 0未读 1已读 2已接收
	buf.PutInt(mail.Time)          //过期时间
	buf.PutString("")              //图片
	buf.PutString(mail.Sender)     //发件人

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

func (p *Proto) MsgGetAttachMail(id int, items [][]int, mailtype int) {
	buf := NewBuffer(1024, p)
	buf.PutUint16(25003) // msgId

	buf.PutUint16(0)
	buf.PutInt(id)
	buf.PutUint16(mailtype)

	buf.PutInt(len(items)) //数量
	if len(items) > 0 {
		for _, e := range items {
			buf.PutInt(e[0])
			buf.PutInt(e[1])
		}
	}

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

func (p *Proto) MsgMailGetCountInfo(unreadCount int, totalCount int) {
	buf := NewBuffer(1024, p)

	buf.PutUint16(25001)
	buf.PutInt(unreadCount)
	buf.PutInt(totalCount)

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

func (p *Proto) MsgBlackMarketDraw(id int, items []int) {
	buf := NewBuffer(1024, p)
	buf.PutUint16(17301) // msgId

	buf.PutUint16(0)
	buf.PutUint16(id)
	buf.PutByte(1)
	buf.PutByte(1)
	buf.PutInt(1)
	buf.PutInt(99)
	buf.PutInt(68)
	buf.PutInt(len(items)) //数量
	if len(items) > 0 {
		for _, e := range items {
			buf.PutInt(1)
			buf.PutInt(e)
			buf.PutInt(1)
			buf.PutByte(0)
		}
	}
	buf.PutUint16(0)
	buf.PutByte(1)

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

func (p *Proto) MsgItemAdd(item *db.HeroItem) {
	buf := NewBuffer(1024, p)
	buf.PutUint16(17004)

	buf.PutInt(item.Id)
	buf.PutByte(1)
	buf.PutByte(0)
	buf.PutInt(item.Id)
	buf.PutInt(item.ItemId)
	buf.PutByte(byte(item.ItemType))
	buf.PutByte(byte(item.IsStoraged))
	buf.PutByte(byte(item.Quality))
	buf.PutByte(byte(item.Star))
	buf.PutByte(byte(item.Lv))
	buf.PutInt(item.Exp)
	buf.PutUint16(item.Role)
	if item.AttrList != "" {
		attrlist := item.GetAttrList()
		buf.PutInt(len(attrlist))
		if len(attrlist) > 0 {
			for _, e := range attrlist {
				buf.PutByte(byte(e[0]))
				buf.PutUint16(e[1])
			}
		}
	} else {
		buf.PutInt(0)
	}
	buf.PutByte(0)   //rindex
	buf.PutUint16(0) //rRole

	buf.PutInt(0)    //rAttrListlength
	buf.PutByte(0)   //isNew
	buf.PutInt(0)    //bigRuneslength
	buf.PutInt(0)    //smallRuneslength
	buf.PutInt(0)    //coreRuneAffixlength
	buf.PutInt(0)    //crysRuneAffixlength
	buf.PutUint16(0) //season

	buf.PutByte(byte(item.IsProtected)) //确定是IsProtected

	buf.PutUint16(0) //RunePunchFailedTime
	buf.PutByte(0)   //是否觉醒

	buf.PutInt(0) //觉醒属性数量 SpecAttrList
	buf.PutInt(0) //StampAttrListlength

	buf.PutInt(0) //宝石孔数量
	buf.PutInt(0) //SpcHoleInfos？

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

func (p *Proto) MsgSmithRecast(code int, itemkey int) {
	buf := NewBuffer(1024, p)
	buf.PutUint16(17007)

	buf.PutUint16(code)
	buf.PutInt(itemkey)

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

func (p *Proto) MsgSmithAttrReplace(code int, itemkey int) {
	buf := NewBuffer(1024, p)
	buf.PutUint16(17008)

	buf.PutUint16(code)
	buf.PutInt(itemkey)

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}

func (p *Proto) MsgSmithAttrUpdate(item *db.HeroItem) {
	buf := NewBuffer(1024, p)
	buf.PutUint16(17012)

	buf.PutInt(item.Id)
	buf.PutInt(item.ItemId)
	buf.PutByte(byte(item.ItemType))
	buf.PutByte(byte(item.IsStoraged))
	buf.PutByte(byte(item.Quality))
	buf.PutByte(byte(item.Star))
	buf.PutByte(byte(item.Lv))
	buf.PutInt(item.Exp)
	buf.PutUint16(item.Role)
	if item.AttrList != "" {
		attrlist := item.GetAttrList()
		buf.PutInt(len(attrlist))
		if len(attrlist) > 0 {
			for _, e := range attrlist {
				buf.PutByte(byte(e[0]))
				buf.PutUint16(e[1])
			}
		}
	} else {
		buf.PutInt(0)
	}
	buf.PutByte(0)   //rindex
	buf.PutUint16(0) //rRole

	if item.RattrList != "" {
		attrlist := item.GetRattrList()
		buf.PutInt(len(attrlist))
		if len(attrlist) > 0 {
			for _, e := range attrlist {
				buf.PutByte(byte(e[0]))
				buf.PutUint16(e[1])
			}
		}
	} else {
		buf.PutInt(0)
	}
	buf.PutByte(0)   //isNew
	buf.PutInt(0)    //bigRuneslength
	buf.PutInt(0)    //smallRuneslength
	buf.PutInt(0)    //coreRuneAffixlength
	buf.PutInt(0)    //crysRuneAffixlength
	buf.PutUint16(0) //season

	buf.PutByte(byte(item.IsProtected)) //确定是IsProtected

	buf.PutUint16(0) //RunePunchFailedTime
	buf.PutByte(0)   //是否觉醒

	buf.PutInt(0) //觉醒属性数量 SpecAttrList
	buf.PutInt(0) //StampAttrListlength

	buf.PutInt(0) //宝石孔数量
	buf.PutInt(0) //SpcHoleInfos？

	buf.Finish(true)
	p.Conn.Write(buf.Byte)
}
