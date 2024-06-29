package role

import (
	"HeroServer/db"
	"fmt"
)

type RoleData struct {
	Role        *db.Role
	Currency    *db.Currency
	CurrencyMap map[int]*db.Currency
	ItemMap     map[int]*db.HeroItem
}

func (r *RoleData) CurrencyUpdate(itemId int, num int) {
	if r.CurrencyMap[itemId] != nil {
		r.CurrencyMap[itemId].Num += num
		db.Conn.Select("num").Save(r.CurrencyMap[itemId])
	} else {
		currency := &db.Currency{
			Rid:    r.Role.Rid,
			ItemId: itemId,
			Num:    num,
		}
		db.Conn.Save(currency)
		fmt.Println(currency.Id)
		if r.CurrencyMap == nil {
			r.CurrencyMap = map[int]*db.Currency{}
		}
		r.CurrencyMap[itemId] = currency
	}
}
