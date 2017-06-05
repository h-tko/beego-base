package models

type UserParam struct {
	Base
	UserID  uint
	Stamina uint
	User    *User `orm:"reverse(one)"`
}
