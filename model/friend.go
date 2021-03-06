package model

import "time"

type Friend struct {
	Id       int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	UserId   int64     `xorm:"bigint(20)" form:"userid" json:"userid"`   // 记录是谁的
	User2Id  int64     `xorm:"bigint(20)" form:"user2id" json:"user2id"` // 对端信息
	CreateAt time.Time `xorm:"datetime" form:"createat" json:"createat"` // 创建时间
}
