package model

type Tabler interface {
	TableName() string
}

func (UserModel) TableName() string {
	return "user"
}
