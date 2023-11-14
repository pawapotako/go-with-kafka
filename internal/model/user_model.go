package model

type UserModel struct {
	Id       uint   `gorm:"column:id;primaryKey;autoIncrement:true"`
	UserId   uint   `gorm:"column:user_id;not null:true"`
	Username string `gorm:"column:username;type:varchar(300);not null:true"`
}
