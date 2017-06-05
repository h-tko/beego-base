package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Base struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func init() {
	orm.RegisterModel(
		new(User),
		new(UserParam),
	)
}
