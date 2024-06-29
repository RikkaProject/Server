package mail

import (
	"HeroServer/db"
	"HeroServer/proto"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func HandleGetAllMail(pt *proto.Proto) {
	var items []*db.Mail
	db.Conn.Where(&db.Mail{Rid: pt.Rid}).Find(&items)
	pt.MPtMailGetAll(items)
}

func HandleGetAttachMail(buf *proto.Buffer, pt *proto.Proto) {
	id := buf.ReadInt()
	var mail db.Mail
	err := db.Conn.Where(&db.Mail{Rid: pt.Rid, Id: id, Status: 0}).First(&mail).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pt.MsgError("邮件不存在或已领取")
		}
		fmt.Println(err)
		return
	}
	var mailType int = 0
	if mail.TemplateId == 1 {
		mailType = 4
	}
	mailItems := mail.GetMailItems()
	if len(mailItems) > 0 && mail.TemplateId == 1 {
		for i, _ := range mailItems {
			pt.RoleData.CurrencyUpdate(mailItems[i][0], mailItems[i][1])
			pt.MsgCurrencyUpdate(mailItems[i][0], uint(pt.RoleData.CurrencyMap[mailItems[i][0]].Num))
			mailItems[i][2] = 1
		}
		mail.SetMailItems(mailItems)
		mail.Status = 2
		db.Conn.Save(mail)
		pt.MsgGetAttachMail(mail.Id, mailItems, mailType)
	}
}

func SendMailCount(pt *proto.Proto) {
	var (
		unread int64
		total  int64
	)
	db.Conn.Model(&db.Mail{}).Where("rid = ? AND status = 0", pt.Rid, 0).Count(&unread)
	db.Conn.Model(&db.Mail{}).Where("rid = ?", pt.Rid).Count(&total)
	fmt.Println(total, unread, "邮件数量")
	pt.MsgMailGetCountInfo(int(unread), int(total))
}
