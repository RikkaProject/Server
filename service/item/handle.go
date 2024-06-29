package item

import (
	"HeroServer/db"
	"HeroServer/gamecfg"
	"HeroServer/proto"
	"math/rand"
	"time"
)

func CreateWeapon(itemid int, pt *proto.Proto) {
	weaponItem := &db.HeroItem{
		Rid:      pt.Rid,
		ItemId:   itemid,
		ItemType: 101,
		Quality:  gamecfg.GameConf.WeaponAttrType[itemid].Quality,
	}
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	attrlen := len(gamecfg.GameConf.WeaponAttrType[itemid].AttrTypes)
	weaponItem.SetAttrList([][]int{
		{gamecfg.GameConf.WeaponAttrType[itemid].AttrTypes[rs.Intn(attrlen)], rs.Intn(901) + 100},
		{gamecfg.GameConf.WeaponAttrType[itemid].AttrTypes[rs.Intn(attrlen)], rs.Intn(901) + 100},
		{gamecfg.GameConf.WeaponAttrType[itemid].AttrTypes[rs.Intn(attrlen)], rs.Intn(901) + 100},
	})

	db.Conn.Create(weaponItem)
	pt.RoleData.ItemMap[weaponItem.Id] = weaponItem

	pt.MsgItemAdd(weaponItem)
}

func WeaponRecast(itemkey int, lockindex []byte, pt *proto.Proto) {
	if pt.RoleData.ItemMap[itemkey] != nil {
		item := pt.RoleData.ItemMap[itemkey]
		rs := rand.New(rand.NewSource(time.Now().UnixNano()))
		attrlen := len(gamecfg.GameConf.WeaponAttrType[item.ItemId].AttrTypes)
		attrlist := item.GetAttrList()
		newatttlist := make([][]int, 3)
		if len(lockindex) != 0 {
			for i := 0; i < 3; i++ {
				for i2 := 0; i2 < len(lockindex); i2++ {
					//fmt.Println((i + 1), int(lockindex[i2]), (i+1) == int(lockindex[i2]))
					if (i + 1) == int(lockindex[i2]) {
						newatttlist[i] = attrlist[i]
						break
					}
				}
				if len(newatttlist[i]) == 0 {
					newatttlist[i] = []int{gamecfg.GameConf.WeaponAttrType[item.ItemId].AttrTypes[rs.Intn(attrlen)], rs.Intn(901) + 100}
				}
			}
		} else {
			for i := 0; i < 3; i++ {
				newatttlist[i] = []int{gamecfg.GameConf.WeaponAttrType[item.ItemId].AttrTypes[rs.Intn(attrlen)], rs.Intn(901) + 100}
			}
		}
		pt.RoleData.ItemMap[itemkey].SetRattrList(newatttlist)

		db.Conn.Save(pt.RoleData.ItemMap[itemkey])

		pt.MsgSmithAttrUpdate(pt.RoleData.ItemMap[itemkey])
		pt.MsgSmithRecast(0, itemkey)

	}
}

func WeaponAttrReplace(itemkey int, isReplace bool, pt *proto.Proto) {
	var code = 0
	if isReplace == true {
		code = 1
		pt.RoleData.ItemMap[itemkey].AttrList = pt.RoleData.ItemMap[itemkey].RattrList
		pt.RoleData.ItemMap[itemkey].RattrList = ""
		db.Conn.Save(pt.RoleData.ItemMap[itemkey])
	} else {
		pt.RoleData.ItemMap[itemkey].RattrList = ""
		db.Conn.Save(pt.RoleData.ItemMap[itemkey])
	}
	pt.MsgSmithAttrUpdate(pt.RoleData.ItemMap[itemkey])
	pt.MsgSmithAttrReplace(code, itemkey)
}
