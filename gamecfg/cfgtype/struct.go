package cfgtype

type LotteryChestCfg struct {
	ID      int `json:"id"`
	Lottery map[string]struct {
		Items []int `json:"items"`
		Prob  int   `json:"prob"`
	}
}

type WeaponAttrTypeCfg struct {
	ID        int   `json:"id"`
	AttrID    int   `json:"attr_id"`
	AttrTypes []int `json:"attr_types"`
	Quality   int   `json:"quality"`
}
