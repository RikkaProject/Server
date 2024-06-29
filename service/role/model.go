package role

import (
	"HeroServer/db"
	"HeroServer/util"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func GetRole(rid int) *db.Role {
	var role db.Role
	err := db.Conn.Where(&db.Role{Rid: rid}).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			role = db.Role{
				Rid:       rid,
				PlayerKey: util.GeneratePlayerKey(rid, 1),
				Name:      "",
				RecordId:  1,
				RoleId:    1,
			}

			// 插入数据
			result := db.Conn.Create(&role)
			if result.Error != nil {
				fmt.Println("创建角色失败:", result.Error)
			} else {
				autoIncrementValue := role.Id
				role.Name = "hero_" + strconv.Itoa(autoIncrementValue)
				db.Conn.Save(&role)
				return &role
			}
		} else {
			return nil
		}
	}
	return &role
}

func CreateRole() {

}
