package lotterychest

import (
	"HeroServer/gamecfg"
	"HeroServer/proto"
	"HeroServer/service/item"
	"fmt"
	"math/rand"
	"time"
)

func HandleOpen(buf *proto.Buffer, pt *proto.Proto) {
	boxid := buf.ReadUint16()
	action := buf.ReadByte()

	if gamecfg.GameConf.LotteryChest[int(boxid)] != nil {
		var prizemap []string
		// 构建奖池
		for s, s2 := range gamecfg.GameConf.LotteryChest[int(boxid)].Lottery {
			for i := 0; i < s2.Prob*10; i++ {
				prizemap = append(prizemap, s)
			}
		}
		rs := rand.New(rand.NewSource(time.Now().UnixNano()))
		// 打乱奖池
		rs.Shuffle(len(prizemap), func(i, j int) {
			prizemap[i], prizemap[j] = prizemap[j], prizemap[i]
		})

		//生成幸运号码
		var lucknum []int
		if action == 1 {
			for i := 0; i < 10; i++ {
				lucknum = append(lucknum, rs.Intn(len(prizemap)))
			}
		} else {
			lucknum = append(lucknum, rs.Intn(len(prizemap)))
		}

		//根据号码获得奖品
		var prize []int
		boxcfg := gamecfg.GameConf.LotteryChest[int(boxid)].Lottery
		for i := 0; i < len(lucknum); i++ {
			switch prizemap[lucknum[i]] {
			case "s":
				prize = append(prize, boxcfg["s"].Items[rs.Intn(len(boxcfg["s"].Items))])
				break
			case "a":
				prize = append(prize, boxcfg["a"].Items[rs.Intn(len(boxcfg["a"].Items))])
				break
			case "b":
				prize = append(prize, boxcfg["b"].Items[rs.Intn(len(boxcfg["b"].Items))])
				break
			case "c":
				prize = append(prize, boxcfg["c"].Items[rs.Intn(len(boxcfg["c"].Items))])
				break
			}
			item.CreateWeapon(prize[i], pt)
		}
		fmt.Println(prize)

		pt.MsgBlackMarketDraw(int(boxid), prize)

	} else {
		pt.MsgError("配置不存在")
	}
}
