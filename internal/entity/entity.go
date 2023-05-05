package entity

type Account struct {
	Id       int64
	UserId   int64
	Service  string
	Login    string
	Password string
}

type User struct {
	Id       int64
	Login    string
	Password string
}
