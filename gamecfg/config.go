package gamecfg

import (
	"HeroServer/gamecfg/cfgtype"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var GameConf *GameConfig

func init() {
	GameConf = &GameConfig{}
	GameConf.LoadCfg()
}

type GameConfig struct {
	ConfDir        string
	LotteryChest   map[int]*cfgtype.LotteryChestCfg
	WeaponAttrType map[int]*cfgtype.WeaponAttrTypeCfg
}

func (cfg *GameConfig) LoadCfg() {
	basedir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取运行目录失败", err)
	}
	cfg.ConfDir = filepath.Join(basedir, "config", "game")
	cfg.loadLotteryChestCfg()
	cfg.loadWeaponAttrTypeCfg()
}

func (cfg *GameConfig) loadLotteryChestCfg() {
	confPath := filepath.Join(cfg.ConfDir, "lotterychest", "chest.json")

	if _, err := os.Stat(confPath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("loadLotteryChestCfg:配置文件不存在")
		}
	}

	confFile, err := os.ReadFile(confPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	var config map[int]*cfgtype.LotteryChestCfg
	err = json.Unmarshal(confFile, &config)
	if err != nil {
		fmt.Println("Error unmarshalling config file:", err)
		return
	}
	cfg.LotteryChest = config
	fmt.Println("loadLotteryChestCfg:加载成功")
	//fmt.Println(config)
}

func (cfg *GameConfig) loadWeaponAttrTypeCfg() {
	confPath := filepath.Join(cfg.ConfDir, "weapon", "weaponAttrType.json")

	if _, err := os.Stat(confPath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("loadWeaponAttrTypeCfg:配置文件不存在")
		}
	}

	confFile, err := os.ReadFile(confPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	var config map[int]*cfgtype.WeaponAttrTypeCfg
	err = json.Unmarshal(confFile, &config)
	if err != nil {
		fmt.Println("Error unmarshalling config file:", err)
		return
	}
	cfg.WeaponAttrType = config
	fmt.Println("loadWeaponAttrTypeCfg:加载成功")
}
