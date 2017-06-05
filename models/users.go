package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Base
	Name        string
	MailAddress string
	Password    string
	UserParam   *UserParam `orm:"rel(one)"`
}

func (user *User) One(id uint) (result *User) {
	o := orm.NewOrm()
	o.QueryTable(user).
		Filter("deleted_at", nil).
		RelatedSel().One(&result)

	return result
}
