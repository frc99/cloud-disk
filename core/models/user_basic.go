package models

type UserBasic struct {
	Id       int
	Identity string
	Name     string
	Password string
	Email    string
}

func (t *UserBasic) tableName() string {
	return "user_basic"
}
