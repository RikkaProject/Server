package db

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Id        int `gorm:"unique;primaryKey;autoIncrement"`
	Rid       int
	PlayerKey uint64 `gorm:"unique"`
	Name      string
	SelfIntro string `gorm:"default:''"`
	RecordId  int
	RoleId    int
	Level     int `gorm:"default:1"`
	Exp       int `gorm:"default:0"`
}

type Currency struct {
	Id     int `gorm:"unique;primaryKey;autoIncrement"`
	Rid    int
	ItemId int
	Num    int `gorm:"default:0"`
}

type Mail struct {
	Id         int `gorm:"unique;primaryKey;autoIncrement"`
	Rid        int
	Status     int
	TemplateId int
	Title      string
	Content    string
	MailItems  string
	Sender     string
	Time       int
}

func (e *Mail) GetMailItems() [][]int {
	var data [][]int
	err := json.Unmarshal([]byte(e.MailItems), &data)
	if err != nil {
		panic(err)
	}
	return data
}

func (e *Mail) SetMailItems(data [][]int) bool {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return false
	}
	e.MailItems = string(jsonStr)
	return false
}

type HeroItem struct {
	Id            int `gorm:"unique;primaryKey;autoIncrement"`
	Rid           int
	ItemId        int
	ItemType      int
	IsStoraged    int    `gorm:"default:0"`
	Quality       int    `gorm:"default:1"`
	Star          int    `gorm:"default:0"`
	Lv            int    `gorm:"default:1"`
	Exp           int    `gorm:"default:0"`
	Role          int    `gorm:"default:0"`
	AttrList      string `gorm:"default:''"`
	Rindex        int    `gorm:"default:0"`
	Rrole         int    `gorm:"default:0"`
	RattrList     string `gorm:"default:''"`
	BigRunes      string `gorm:"default:''"`
	SmallRunes    string `gorm:"default:''"`
	CoreRuneAffix string `gorm:"default:''"`
	CrysRuneAffix string `gorm:"default:''"`
	Season        int    `gorm:"default:0"`
	IsProtected   int    `gorm:"default:0"`
	IsAwake       int    `gorm:"default:0"`
	SpecAttrList  string `gorm:"default:''"`
	StampAttrList string `gorm:"default:''"`
	GemList       string `gorm:"default:''"`
}

func (e *HeroItem) SetAttrList(data [][]int) bool {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return false
	}
	e.AttrList = string(jsonStr)
	return false
}

func (e *HeroItem) GetAttrList() [][]int {
	var data [][]int
	err := json.Unmarshal([]byte(e.AttrList), &data)
	if err != nil {
		panic(err)
	}
	return data
}

func (e *HeroItem) SetRattrList(data [][]int) bool {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return false
	}
	e.RattrList = string(jsonStr)
	return false
}

func (e *HeroItem) GetRattrList() [][]int {
	var data [][]int
	err := json.Unmarshal([]byte(e.RattrList), &data)
	if err != nil {
		panic(err)
	}
	return data
}
